package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/relentlessworks/casekit/internal/model"
)

// Handler holds all HTTP handlers for the casekit service.
type Handler struct {
	secret string
}

// New creates a new API handler.
func New(secret string) *Handler {
	return &Handler{secret: secret}
}

// Routes returns the HTTP mux with all routes registered.
func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/help", h.help)
	mux.HandleFunc("/.well-known/agent.md", h.help)
	mux.HandleFunc("/convert", h.convert)
	mux.HandleFunc("/detect", h.detect)
	mux.HandleFunc("/all", h.all)
	mux.HandleFunc("/camel", h.camel)
	mux.HandleFunc("/pascal", h.pascal)
	mux.HandleFunc("/snake", h.snake)
	mux.HandleFunc("/screaming-snake", h.screamingSnake)
	mux.HandleFunc("/kebab", h.kebab)
	mux.HandleFunc("/screaming-kebab", h.screamingKebab)
	mux.HandleFunc("/dot", h.dot)
	mux.HandleFunc("/title", h.title)
	mux.HandleFunc("/sentence", h.sentence)
	mux.HandleFunc("/flat", h.flat)
	mux.HandleFunc("/train", h.train)
	mux.HandleFunc("/alternating", h.alternating)
	mux.HandleFunc("/inverse", h.inverse)
	mux.HandleFunc("/mcp", h.mcp)
	return mux
}

// wantsJSON checks if the client wants JSON output.
func wantsJSON(r *http.Request) bool {
	if r.URL.Query().Get("format") == "json" {
		return true
	}
	return strings.Contains(r.Header.Get("Accept"), "application/json")
}

// writeError writes an instructive error response.
func writeError(w http.ResponseWriter, status int, msg, hint string) {
	w.WriteHeader(status)
	if hint != "" {
		fmt.Fprintf(w, "error: %s | hint: %s\n", msg, hint)
	} else {
		fmt.Fprintf(w, "error: %s\n", msg)
	}
}

// writeJSON writes a JSON response.
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// getTextParam extracts the text from query params or form body.
func getTextParam(r *http.Request) (string, error) {
	val := r.URL.Query().Get("text")
	if val == "" {
		val = r.FormValue("text")
	}
	if val == "" {
		return "", fmt.Errorf("missing text parameter")
	}
	return val, nil
}

// parseCaseType parses a case type string, normalizing common aliases.
func parseCaseType(s string) (model.CaseType, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	switch s {
	case "camel", "camelcase":
		return model.CaseCamel, nil
	case "pascal", "pascalcase":
		return model.CasePascal, nil
	case "snake", "snake_case":
		return model.CaseSnake, nil
	case "screaming_snake", "screaming-snake", "screamingsnake", "constant", "constant_case":
		return model.CaseScreamingSnake, nil
	case "kebab", "kebab-case":
		return model.CaseKebab, nil
	case "screaming_kebab", "screaming-kebab", "screamingkebab":
		return model.CaseScreamingKebab, nil
	case "dot", "dotcase":
		return model.CaseDot, nil
	case "title", "titlecase", "title_case":
		return model.CaseTitle, nil
	case "sentence", "sentencecase", "sentence_case":
		return model.CaseSentence, nil
	case "flat", "flatcase":
		return model.CaseFlat, nil
	case "train", "traincase", "train_case":
		return model.CaseTrain, nil
	case "alternating", "alternatingcase":
		return model.CaseAlternating, nil
	case "inverse", "inversecase", "swap", "swapcase":
		return model.CaseInverse, nil
	default:
		return "", fmt.Errorf("unknown case type: %s", s)
	}
}

// writeResult writes a single conversion result in plain text or JSON.
func writeResult(w http.ResponseWriter, r *http.Request, caseType string, result string) {
	if wantsJSON(r) {
		writeJSON(w, map[string]string{caseType: result})
		return
	}
	fmt.Fprintf(w, "%s=%s\n", caseType, result)
}

// writeConversion writes a conversion result with the case name as label.
func writeConversion(w http.ResponseWriter, r *http.Request, ct model.CaseType, text string) {
	result := model.Convert(text, ct)
	if wantsJSON(r) {
		writeJSON(w, map[string]string{string(ct): result})
		return
	}
	fmt.Fprintf(w, "%s=%s\n", ct, result)
}

// --- Handlers ---

func (h *Handler) help(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, helpText)
}

func (h *Handler) convert(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=helloWorld&to=snake")
		return
	}
	toStr := r.URL.Query().Get("to")
	if toStr == "" {
		toStr = r.FormValue("to")
	}
	if toStr == "" {
		writeError(w, http.StatusBadRequest, "missing to parameter", "specify target case: camel, pascal, snake, screaming_snake, kebab, screaming_kebab, dot, title, sentence, flat, train, alternating, inverse")
		return
	}
	ct, err := parseCaseType(toStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "valid cases: camel, pascal, snake, screaming_snake, kebab, screaming_kebab, dot, title, sentence, flat, train, alternating, inverse")
		return
	}
	result := model.Convert(text, ct)
	if wantsJSON(r) {
		writeJSON(w, map[string]interface{}{
			"input":  text,
			"to":     string(ct),
			"result": result,
		})
		return
	}
	fmt.Fprintf(w, "%s=%s\n", ct, result)
}

func (h *Handler) detect(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=helloWorld")
		return
	}
	detected := model.Detect(text)
	name := model.CaseName(detected)
	if wantsJSON(r) {
		writeJSON(w, map[string]string{
			"detected": string(detected),
			"name":     name,
		})
		return
	}
	fmt.Fprintf(w, "detected=%s name=%s\n", detected, name)
}

func (h *Handler) all(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=helloWorld")
		return
	}
	conversions := model.AllConversions(text)
	if wantsJSON(r) {
		result := make(map[string]string)
		for ct, val := range conversions {
			result[string(ct)] = val
		}
		writeJSON(w, result)
		return
	}
	for _, ct := range model.AllCases() {
		fmt.Fprintf(w, "%s=%s\n", ct, conversions[ct])
	}
}

func (h *Handler) camel(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=hello_world")
		return
	}
	writeConversion(w, r, model.CaseCamel, text)
}

func (h *Handler) pascal(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=hello_world")
		return
	}
	writeConversion(w, r, model.CasePascal, text)
}

func (h *Handler) snake(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=helloWorld")
		return
	}
	writeConversion(w, r, model.CaseSnake, text)
}

func (h *Handler) screamingSnake(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=helloWorld")
		return
	}
	writeConversion(w, r, model.CaseScreamingSnake, text)
}

func (h *Handler) kebab(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=helloWorld")
		return
	}
	writeConversion(w, r, model.CaseKebab, text)
}

func (h *Handler) screamingKebab(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=helloWorld")
		return
	}
	writeConversion(w, r, model.CaseScreamingKebab, text)
}

func (h *Handler) dot(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=helloWorld")
		return
	}
	writeConversion(w, r, model.CaseDot, text)
}

func (h *Handler) title(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=hello_world")
		return
	}
	writeConversion(w, r, model.CaseTitle, text)
}

func (h *Handler) sentence(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=hello_world")
		return
	}
	writeConversion(w, r, model.CaseSentence, text)
}

func (h *Handler) flat(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=hello_world")
		return
	}
	writeConversion(w, r, model.CaseFlat, text)
}

func (h *Handler) train(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=hello_world")
		return
	}
	writeConversion(w, r, model.CaseTrain, text)
}

func (h *Handler) alternating(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=hello")
		return
	}
	writeConversion(w, r, model.CaseAlternating, text)
}

func (h *Handler) inverse(w http.ResponseWriter, r *http.Request) {
	text, err := getTextParam(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error(), "provide text as ?text=HelloWorld")
		return
	}
	writeConversion(w, r, model.CaseInverse, text)
}
