package main

import (
	"encoding/hex"
	"syscall/js"
)

// State represents the current state of AES encryption/decryption
type State struct {
	Data         [4][4]byte
	Key          [4][4]byte
	CurrentStep  string
	CurrentRound int
}

// AESStep performs a single step of AES encryption visualization
func AESStep(state *State, step string) {
	state.CurrentStep = step

	switch step {
	case "SubBytes":
		// Substitute each byte using the S-box (simplified for visualization)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				// Simple substitution for visualization purposes
				state.Data[i][j] = subByte(state.Data[i][j])
			}
		}
	case "ShiftRows":
		// Shift rows by different offsets
		// Row 0: shift by 0
		// Row 1: shift by 1
		// Row 2: shift by 2
		// Row 3: shift by 3
		for i := 1; i < 4; i++ {
			shiftRow(state, i, i)
		}
	case "MixColumns":
		// Simple column mixing for visualization (not the actual complex operation)
		mixColumns(state)
	case "AddRoundKey":
		// XOR the state with the round key
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				state.Data[i][j] ^= state.Key[i][j]
			}
		}
	}
}

// Simple S-box substitute for visualization
func subByte(b byte) byte {
	// This is a simplified substitution, not the actual AES S-box
	return (b << 1) ^ (b >> 7) ^ 0x63
}

// Shift a row in the state by a specified offset
func shiftRow(state *State, row, shift int) {
	// Create a copy of the row
	var temp [4]byte
	for i := 0; i < 4; i++ {
		temp[i] = state.Data[row][i]
	}

	// Shift the row
	for i := 0; i < 4; i++ {
		state.Data[row][i] = temp[(i+shift)%4]
	}
}

// Simplified column mixing (not the actual complex matrix multiplication)
func mixColumns(state *State) {
	for col := 0; col < 4; col++ {
		a := state.Data[0][col]
		b := state.Data[1][col]
		c := state.Data[2][col]
		d := state.Data[3][col]

		// Simple mix for visualization
		state.Data[0][col] = a ^ b
		state.Data[1][col] = b ^ c
		state.Data[2][col] = c ^ d
		state.Data[3][col] = d ^ a
	}
}

// InitState creates a new AES state from text and key
func InitState(text, key string) *State {
	state := &State{
		CurrentStep:  "Initial",
		CurrentRound: 0,
	}

	// Fill state with text data (simplified)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i*4+j < len(text) {
				state.Data[j][i] = byte(text[i*4+j])
			}
		}
	}

	// Fill key (simplified)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i*4+j < len(key) {
				state.Key[j][i] = byte(key[i*4+j])
			}
		}
	}

	return state
}

// StateToHex converts the current state to a hex string for display
func StateToHex(state *State) string {
	result := ""
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result += hex.EncodeToString([]byte{state.Data[j][i]}) + " "
		}
		result += "\n"
	}
	return result
}

// JS wrapper for InitState
func jsInitState(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Invalid number of arguments"
	}

	text := args[0].String()
	key := args[1].String()

	state := InitState(text, key)

	// Convert state to JS object
	result := make(map[string]interface{})
	result["data"] = stateDataToJSArray(state.Data)
	result["key"] = stateDataToJSArray(state.Key)
	result["step"] = state.CurrentStep
	result["round"] = state.CurrentRound

	return js.ValueOf(result)
}

// JS wrapper for AESStep
func jsAESStep(this js.Value, args []js.Value) interface{} {
	if len(args) != 5 {
		return "Invalid number of arguments"
	}

	// Extract state from JS arguments
	state := &State{
		CurrentStep:  args[2].String(),
		CurrentRound: args[3].Int(),
	}

	// Extract data array
	dataObj := args[0]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			state.Data[i][j] = byte(dataObj.Index(i).Index(j).Int())
		}
	}

	// Extract key array
	keyObj := args[1]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			state.Key[i][j] = byte(keyObj.Index(i).Index(j).Int())
		}
	}

	// Perform the AES step
	step := args[4].String()
	AESStep(state, step)

	// Convert state back to JS object
	result := make(map[string]interface{})
	result["data"] = stateDataToJSArray(state.Data)
	result["key"] = stateDataToJSArray(state.Key)
	result["step"] = state.CurrentStep
	result["round"] = state.CurrentRound

	return js.ValueOf(result)
}

// Helper function to convert state data to JS array
func stateDataToJSArray(data [4][4]byte) interface{} {
	result := make([]interface{}, 4)
	for i := 0; i < 4; i++ {
		row := make([]interface{}, 4)
		for j := 0; j < 4; j++ {
			row[j] = data[i][j]
		}
		result[i] = row
	}
	return result
}

func main() {
	c := make(chan struct{})

	// Register functions to be called from JavaScript
	js.Global().Set("initAESState", js.FuncOf(jsInitState))
	js.Global().Set("performAESStep", js.FuncOf(jsAESStep))

	// Keep the program running
	<-c
}
