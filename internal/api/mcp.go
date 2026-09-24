package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/relentlessworks/casekit/internal/model"
)

type mcpRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type mcpResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *mcpError   `json:"error,omitempty"`
}

type mcpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *Handler) mcp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed", "POST JSON-RPC 2.0 to /mcp")
		return
	}

	var req mcpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeMCPError(w, nil, -32700, "parse error")
		return
	}

	switch req.Method {
	case "initialize":
		writeMCPResult(w, req.ID, map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"serverInfo": map[string]string{
				"name":    "casekit",
				"version": "0.1.0",
			},
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
		})

	case "tools/list":
		writeMCPResult(w, req.ID, map[string]interface{}{
			"tools": mcpTools(),
		})

	case "tools/call":
		var params struct {
			Name      string            `json:"name"`
			Arguments map[string]string `json:"arguments"`
		}
		if err := json.Unmarshal(req.Params, &params); err != nil {
			writeMCPError(w, req.ID, -32602, "invalid params")
			return
		}
		result, err := h.handleMCPTool(params.Name, params.Arguments)
		if err != nil {
			writeMCPError(w, req.ID, -32603, err.Error())
			return
		}
		writeMCPResult(w, req.ID, map[string]interface{}{
			"content": []map[string]string{
				{"type": "text", "text": result},
			},
		})

	default:
		writeMCPError(w, req.ID, -32601, "method not found")
	}
}

func writeMCPResult(w http.ResponseWriter, id interface{}, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mcpResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	})
}

func writeMCPError(w http.ResponseWriter, id interface{}, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mcpResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error:   &mcpError{Code: code, Message: msg},
	})
}

type mcpTool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema map[string]interface{} `json:"inputSchema"`
}

func mcpTools() []mcpTool {
	strType := map[string]interface{}{"type": "string"}
	textProp := map[string]interface{}{"type": "string", "description": "the text to convert"}
	toProp := map[string]interface{}{"type": "string", "description": "target case: camel, pascal, snake, screaming_snake, kebab, screaming_kebab, dot, title, sentence, flat, train, alternating, inverse"}

	tools := []mcpTool{
		{Name: "convert", Description: "Convert text to a specific case format", InputSchema: schemaProps(map[string]interface{}{"text": textProp, "to": toProp}, "text", "to")},
		{Name: "detect", Description: "Detect the case type of text", InputSchema: schemaProps(map[string]interface{}{"text": textProp}, "text")},
		{Name: "all", Description: "Convert text to all supported case formats", InputSchema: schemaProps(map[string]interface{}{"text": textProp}, "text")},
	}

	for _, ct := range model.AllCases() {
		name := string(ct)
		desc := fmt.Sprintf("Convert text to %s", model.CaseName(ct))
		tools = append(tools, mcpTool{
			Name:        name,
			Description: desc,
			InputSchema: schemaProps(map[string]interface{}{"text": textProp}, "text"),
		})
	}

	_ = strType
	return tools
}

func schemaProps(props map[string]interface{}, required ...string) map[string]interface{} {
	return map[string]interface{}{
		"type":       "object",
		"properties": props,
		"required":   required,
	}
}

func (h *Handler) handleMCPTool(name string, args map[string]string) (string, error) {
	text, ok := args["text"]
	if !ok || text == "" {
		return "", fmt.Errorf("missing text argument")
	}

	switch name {
	case "convert":
		toStr, ok := args["to"]
		if !ok || toStr == "" {
			return "", fmt.Errorf("missing to argument")
		}
		ct, err := parseCaseType(toStr)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s=%s", ct, model.Convert(text, ct)), nil

	case "detect":
		detected := model.Detect(text)
		return fmt.Sprintf("detected=%s name=%s", detected, model.CaseName(detected)), nil

	case "all":
		conversions := model.AllConversions(text)
		var sb strings.Builder
		for _, ct := range model.AllCases() {
			sb.WriteString(fmt.Sprintf("%s=%s\n", ct, conversions[ct]))
		}
		return sb.String(), nil

	default:
		ct, err := parseCaseType(name)
		if err != nil {
			return "", fmt.Errorf("unknown tool: %s", name)
		}
		return fmt.Sprintf("%s=%s", ct, model.Convert(text, ct)), nil
	}
}
