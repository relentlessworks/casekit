package api

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/convert?text=hello_world&to=camel", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("convert status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "camel=helloWorld") {
		t.Errorf("convert body = %q, want camel=helloWorld", body)
	}
}

func TestConvertJSON(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/convert?text=hello_world&to=snake&format=json", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("convert status = %d, want 200", w.Code)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("convert JSON parse error: %v", err)
	}
	if result["result"] != "hello_world" {
		t.Errorf("convert JSON result = %v, want hello_world", result["result"])
	}
}

func TestConvertMissingText(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/convert?to=camel", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("convert missing text status = %d, want 400", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "error:") {
		t.Errorf("convert missing text body = %q, want error", body)
	}
	if !strings.Contains(body, "hint:") {
		t.Errorf("convert missing text body = %q, want hint", body)
	}
}

func TestConvertMissingTo(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/convert?text=hello", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("convert missing to status = %d, want 400", w.Code)
	}
}

func TestConvertInvalidCase(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/convert?text=hello&to=foobar", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("convert invalid case status = %d, want 400", w.Code)
	}
}

func TestDetect(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/detect?text=helloWorld", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("detect status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "detected=camel") {
		t.Errorf("detect body = %q, want detected=camel", body)
	}
}

func TestDetectJSON(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/detect?text=hello_world&format=json", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("detect status = %d, want 200", w.Code)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("detect JSON parse error: %v", err)
	}
	if result["detected"] != "snake" {
		t.Errorf("detect JSON detected = %v, want snake", result["detected"])
	}
}

func TestAll(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/all?text=helloWorld", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("all status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	lines := strings.Split(strings.TrimSpace(body), "\n")
	if len(lines) < 13 {
		t.Errorf("all lines = %d, want >= 13", len(lines))
	}
	if !strings.Contains(body, "camel=helloWorld") {
		t.Errorf("all body should contain camel=helloWorld, got %q", body)
	}
	if !strings.Contains(body, "snake=hello_world") {
		t.Errorf("all body should contain snake=hello_world, got %q", body)
	}
}

func TestAllJSON(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/all?text=hello&format=json", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("all status = %d, want 200", w.Code)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("all JSON parse error: %v", err)
	}
	if result["camel"] != "hello" {
		t.Errorf("all JSON camel = %v, want hello", result["camel"])
	}
}

func TestCamel(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/camel?text=hello_world", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("camel status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "camel=helloWorld") {
		t.Errorf("camel body = %q, want camel=helloWorld", body)
	}
}

func TestPascal(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/pascal?text=hello_world", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("pascal status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "pascal=HelloWorld") {
		t.Errorf("pascal body = %q, want pascal=HelloWorld", body)
	}
}

func TestSnake(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/snake?text=helloWorld", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("snake status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "snake=hello_world") {
		t.Errorf("snake body = %q, want snake=hello_world", body)
	}
}

func TestScreamingSnake(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/screaming-snake?text=helloWorld", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("screaming-snake status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "screaming_snake=HELLO_WORLD") {
		t.Errorf("screaming-snake body = %q, want screaming_snake=HELLO_WORLD", body)
	}
}

func TestKebab(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/kebab?text=helloWorld", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("kebab status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "kebab=hello-world") {
		t.Errorf("kebab body = %q, want kebab=hello-world", body)
	}
}

func TestScreamingKebab(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/screaming-kebab?text=helloWorld", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("screaming-kebab status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "screaming_kebab=HELLO-WORLD") {
		t.Errorf("screaming-kebab body = %q, want screaming_kebab=HELLO-WORLD", body)
	}
}

func TestDot(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/dot?text=helloWorld", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("dot status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "dot=hello.world") {
		t.Errorf("dot body = %q, want dot=hello.world", body)
	}
}

func TestTitle(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/title?text=hello_world", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("title status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "title=Hello World") {
		t.Errorf("title body = %q, want title=Hello World", body)
	}
}

func TestSentence(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/sentence?text=hello_world", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("sentence status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "sentence=Hello world") {
		t.Errorf("sentence body = %q, want sentence=Hello world", body)
	}
}

func TestFlat(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/flat?text=hello_world", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("flat status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "flat=helloworld") {
		t.Errorf("flat body = %q, want flat=helloworld", body)
	}
}

func TestTrain(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/train?text=hello_world", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("train status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "train=Hello-World") {
		t.Errorf("train body = %q, want train=Hello-World", body)
	}
}

func TestAlternating(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/alternating?text=hello", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("alternating status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "alternating=hElLo") {
		t.Errorf("alternating body = %q, want alternating=hElLo", body)
	}
}

func TestInverse(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/inverse?text=HelloWorld", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("inverse status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "inverse=hELLOwORLD") {
		t.Errorf("inverse body = %q, want inverse=hELLOwORLD", body)
	}
}

func TestHelp(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("GET", "/help", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("help status = %d, want 200", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "casekit") {
		t.Errorf("help body should contain 'casekit'")
	}
	if !strings.Contains(body, "ENDPOINTS") {
		t.Errorf("help body should contain 'ENDPOINTS'")
	}
}

func TestMCPInitialize(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader(`{"jsonrpc":"2.0","id":1,"method":"initialize"}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("mcp initialize status = %d, want 200", w.Code)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("mcp initialize JSON parse error: %v", err)
	}
	if result["jsonrpc"] != "2.0" {
		t.Errorf("mcp initialize jsonrpc = %v, want 2.0", result["jsonrpc"])
	}
}

func TestMCPToolsList(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader(`{"jsonrpc":"2.0","id":2,"method":"tools/list"}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("mcp tools/list status = %d, want 200", w.Code)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("mcp tools/list JSON parse error: %v", err)
	}
	tools, ok := result["result"].(map[string]interface{})["tools"].([]interface{})
	if !ok {
		t.Errorf("mcp tools/list result not as expected")
		return
	}
	if len(tools) < 16 {
		t.Errorf("mcp tools/list count = %d, want >= 16", len(tools))
	}
}

func TestMCPToolsCall(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader(`{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"convert","arguments":{"text":"hello_world","to":"camel"}}}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("mcp tools/call status = %d, want 200", w.Code)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("mcp tools/call JSON parse error: %v", err)
	}
	content, ok := result["result"].(map[string]interface{})["content"].([]interface{})
	if !ok {
		t.Errorf("mcp tools/call result not as expected")
		return
	}
	text := content[0].(map[string]interface{})["text"].(string)
	if !strings.Contains(text, "camel=helloWorld") {
		t.Errorf("mcp tools/call text = %q, want camel=helloWorld", text)
	}
}

func TestMCPToolsCallDetect(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader(`{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"detect","arguments":{"text":"helloWorld"}}}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("mcp tools/call detect status = %d, want 200", w.Code)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("mcp tools/call detect JSON parse error: %v", err)
	}
	content, ok := result["result"].(map[string]interface{})["content"].([]interface{})
	if !ok {
		t.Errorf("mcp tools/call detect result not as expected")
		return
	}
	text := content[0].(map[string]interface{})["text"].(string)
	if !strings.Contains(text, "detected=camel") {
		t.Errorf("mcp tools/call detect text = %q, want detected=camel", text)
	}
}

func TestMCPToolsCallAll(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader(`{"jsonrpc":"2.0","id":5,"method":"tools/call","params":{"name":"all","arguments":{"text":"helloWorld"}}}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("mcp tools/call all status = %d, want 200", w.Code)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("mcp tools/call all JSON parse error: %v", err)
	}
	content, ok := result["result"].(map[string]interface{})["content"].([]interface{})
	if !ok {
		t.Errorf("mcp tools/call all result not as expected")
		return
	}
	text := content[0].(map[string]interface{})["text"].(string)
	if !strings.Contains(text, "camel=helloWorld") {
		t.Errorf("mcp tools/call all text = %q, want camel=helloWorld", text)
	}
}

func TestMCPToolsCallIndividual(t *testing.T) {
	h := New("test-secret")
	mux := h.Routes()

	req := httptest.NewRequest("POST", "/mcp", strings.NewReader(`{"jsonrpc":"2.0","id":6,"method":"tools/call","params":{"name":"snake","arguments":{"text":"helloWorld"}}}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("mcp tools/call snake status = %d, want 200", w.Code)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Errorf("mcp tools/call snake JSON parse error: %v", err)
	}
	content, ok := result["result"].(map[string]interface{})["content"].([]interface{})
	if !ok {
		t.Errorf("mcp tools/call snake result not as expected")
		return
	}
	text := content[0].(map[string]interface{})["text"].(string)
	if !strings.Contains(text, "snake=hello_world") {
		t.Errorf("mcp tools/call snake text = %q, want snake=hello_world", text)
	}
}
