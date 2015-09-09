// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Input commands.
// API Version: 1.1

package gcd

import (
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Input() *ChromeInput {
	if c.input == nil {
		c.input = newChromeInput(c)
	}
	return c.input
}

type ChromeInput struct {
	target *ChromeTarget
}

func newChromeInput(target *ChromeTarget) *ChromeInput {
	c := &ChromeInput{target: target}
	return c
}

// dispatchKeyEvent - Dispatches a key event to the page.
// type - Type of the key event.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time).
// text - Text as generated by processing a virtual key code with a keyboard layout. Not needed for for <code>keyUp</code> and <code>rawKeyDown</code> events (default: "")
// unmodifiedText - Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
// keyIdentifier - Unique key identifier (e.g., 'U+0041') (default: "").
// code - Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
// key - Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
// windowsVirtualKeyCode - Windows virtual key code (default: 0).
// nativeVirtualKeyCode - Native virtual key code (default: 0).
// autoRepeat - Whether the event was generated from auto repeat (default: false).
// isKeypad - Whether the event was generated from the keypad (default: false).
// isSystemKey - Whether the event was a system key event (default: false).
func (c *ChromeInput) DispatchKeyEvent(theType string, modifiers int, timestamp float64, text string, unmodifiedText string, keyIdentifier string, code string, key string, windowsVirtualKeyCode int, nativeVirtualKeyCode int, autoRepeat bool, isKeypad bool, isSystemKey bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 13)
	paramRequest["type"] = theType
	paramRequest["modifiers"] = modifiers
	paramRequest["timestamp"] = timestamp
	paramRequest["text"] = text
	paramRequest["unmodifiedText"] = unmodifiedText
	paramRequest["keyIdentifier"] = keyIdentifier
	paramRequest["code"] = code
	paramRequest["key"] = key
	paramRequest["windowsVirtualKeyCode"] = windowsVirtualKeyCode
	paramRequest["nativeVirtualKeyCode"] = nativeVirtualKeyCode
	paramRequest["autoRepeat"] = autoRepeat
	paramRequest["isKeypad"] = isKeypad
	paramRequest["isSystemKey"] = isSystemKey
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.dispatchKeyEvent", Params: paramRequest})
}

// dispatchMouseEvent - Dispatches a mouse event to the page.
// type - Type of the mouse event.
// x - X coordinate of the event relative to the main frame's viewport.
// y - Y coordinate of the event relative to the main frame's viewport. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time).
// button - Mouse button (default: "none").
// clickCount - Number of times the mouse button was clicked (default: 0).
func (c *ChromeInput) DispatchMouseEvent(theType string, x int, y int, modifiers int, timestamp float64, button string, clickCount int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 7)
	paramRequest["type"] = theType
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["modifiers"] = modifiers
	paramRequest["timestamp"] = timestamp
	paramRequest["button"] = button
	paramRequest["clickCount"] = clickCount
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.dispatchMouseEvent", Params: paramRequest})
}

// dispatchTouchEvent - Dispatches a touch event to the page.
// type - Type of the touch event.
// touchPoints - Touch points.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time).
func (c *ChromeInput) DispatchTouchEvent(theType string, touchPoints []*types.ChromeInputTouchPoint, modifiers int, timestamp float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["type"] = theType
	paramRequest["touchPoints"] = touchPoints
	paramRequest["modifiers"] = modifiers
	paramRequest["timestamp"] = timestamp
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.dispatchTouchEvent", Params: paramRequest})
}

// emulateTouchFromMouseEvent - Emulates touch event from the mouse event parameters.
// type - Type of the mouse event.
// x - X coordinate of the mouse pointer in DIP.
// y - Y coordinate of the mouse pointer in DIP.
// timestamp - Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970.
// button - Mouse button.
// deltaX - X delta in DIP for mouse wheel event (default: 0).
// deltaY - Y delta in DIP for mouse wheel event (default: 0).
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// clickCount - Number of times the mouse button was clicked (default: 0).
func (c *ChromeInput) EmulateTouchFromMouseEvent(theType string, x int, y int, timestamp float64, button string, deltaX float64, deltaY float64, modifiers int, clickCount int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 9)
	paramRequest["type"] = theType
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["timestamp"] = timestamp
	paramRequest["button"] = button
	paramRequest["deltaX"] = deltaX
	paramRequest["deltaY"] = deltaY
	paramRequest["modifiers"] = modifiers
	paramRequest["clickCount"] = clickCount
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.emulateTouchFromMouseEvent", Params: paramRequest})
}

// synthesizePinchGesture - Synthesizes a pinch gesture over a time period by issuing appropriate touch events.
// x - X coordinate of the start of the gesture in CSS pixels.
// y - Y coordinate of the start of the gesture in CSS pixels.
// scaleFactor - Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
// relativeSpeed - Relative pointer speed in pixels per second (default: 800).
// gestureSourceType - Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
func (c *ChromeInput) SynthesizePinchGesture(x int, y int, scaleFactor float64, relativeSpeed int, gestureSourceType *types.ChromeInputGestureSourceType) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["scaleFactor"] = scaleFactor
	paramRequest["relativeSpeed"] = relativeSpeed
	paramRequest["gestureSourceType"] = gestureSourceType
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.synthesizePinchGesture", Params: paramRequest})
}

// synthesizeScrollGesture - Synthesizes a scroll gesture over a time period by issuing appropriate touch events.
// x - X coordinate of the start of the gesture in CSS pixels.
// y - Y coordinate of the start of the gesture in CSS pixels.
// xDistance - The distance to scroll along the X axis (positive to scroll left).
// yDistance - The distance to scroll along the Y axis (positive to scroll up).
// xOverscroll - The number of additional pixels to scroll back along the X axis, in addition to the given distance.
// yOverscroll - The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
// preventFling - Prevent fling (default: true).
// speed - Swipe speed in pixels per second (default: 800).
// gestureSourceType - Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
// repeatCount - The number of times to repeat the gesture (default: 0).
// repeatDelayMs - The number of milliseconds delay between each repeat. (default: 250).
// interactionMarkerName - The name of the interaction markers to generate, if not empty (default: "").
func (c *ChromeInput) SynthesizeScrollGesture(x int, y int, xDistance int, yDistance int, xOverscroll int, yOverscroll int, preventFling bool, speed int, gestureSourceType *types.ChromeInputGestureSourceType, repeatCount int, repeatDelayMs int, interactionMarkerName string) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 12)
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["xDistance"] = xDistance
	paramRequest["yDistance"] = yDistance
	paramRequest["xOverscroll"] = xOverscroll
	paramRequest["yOverscroll"] = yOverscroll
	paramRequest["preventFling"] = preventFling
	paramRequest["speed"] = speed
	paramRequest["gestureSourceType"] = gestureSourceType
	paramRequest["repeatCount"] = repeatCount
	paramRequest["repeatDelayMs"] = repeatDelayMs
	paramRequest["interactionMarkerName"] = interactionMarkerName
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.synthesizeScrollGesture", Params: paramRequest})
}

// synthesizeTapGesture - Synthesizes a tap gesture over a time period by issuing appropriate touch events.
// x - X coordinate of the start of the gesture in CSS pixels.
// y - Y coordinate of the start of the gesture in CSS pixels.
// duration - Duration between touchdown and touchup events in ms (default: 50).
// tapCount - Number of times to perform the tap (e.g. 2 for double tap, default: 1).
// gestureSourceType - Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
func (c *ChromeInput) SynthesizeTapGesture(x int, y int, duration int, tapCount int, gestureSourceType *types.ChromeInputGestureSourceType) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 5)
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["duration"] = duration
	paramRequest["tapCount"] = tapCount
	paramRequest["gestureSourceType"] = gestureSourceType
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.synthesizeTapGesture", Params: paramRequest})
}
