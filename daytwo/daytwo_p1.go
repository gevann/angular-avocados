package daytwo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OpCode uint8

const (
	forward OpCode = iota
	down
	up
)

type Instruction struct {
	opCode OpCode
	value  uint8
}

type Result[T any] struct {
	Data  T
	Error error
}

func parseOpCode(s string) (OpCode, error) {
	var zeroValue OpCode
	// switch on the first character
	switch s[0] {
	case 'f':
		return forward, nil
	case 'd':
		return down, nil
	case 'u':
		return up, nil
	case 'F':
		return forward, nil
	case 'D':
		return down, nil
	case 'U':
		return up, nil
	default:
		return zeroValue, fmt.Errorf("Unknown opcode: %s", s)
	}
}

func parse(s string) (Instruction, error) {
	split := strings.Split(s, " ")
	if len(split) != 2 {
		return Instruction{}, fmt.Errorf("Invalid instruction: %s", s)
	}
	opCode, err := parseOpCode(split[0])
	if err != nil {
		return Instruction{}, err
	}
	value, err := strconv.Atoi(split[1])
	if err != nil {
		return Instruction{}, err
	}
	if value < 0 {
		return Instruction{}, fmt.Errorf("Invalid value: %d", value)
	}

	return Instruction{
		opCode: opCode,
		value:  uint8(value),
	}, nil
}

func channelInstructions(filename string) (<-chan Result[Instruction], error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	ch := make(chan Result[Instruction])

	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			instruction, err := parse(line)
			if err != nil {
				ch <- Result[Instruction]{Error: err}
			} else {
				ch <- Result[Instruction]{Data: instruction}
			}
		}
	}()

	return ch, nil
}

func PartOne(filename string) (int, error) {
	ch, err := channelInstructions(filename)
	if err != nil {
		return 0, err
	}

	var horizontal, depth int

	for instruction := range ch {
		if instruction.Error != nil {
			return 0, instruction.Error
		}
		switch instruction.Data.opCode {
		case forward:
			horizontal += int(instruction.Data.value)
		case down:
			depth += int(instruction.Data.value)
		case up:
			depth -= int(instruction.Data.value)
		default:
			return 0, fmt.Errorf("Unknown opcode: %d", instruction.Data.opCode)
		}
	}

	product := horizontal * depth

	return product, nil
}
