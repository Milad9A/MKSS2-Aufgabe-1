package main

import "time"

// Position represents the robot's coordinates
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Action represents an activity performed by a robot
type Action struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Details   string    `json:"details"`
}

// Robot represents a robot in the system
type Robot struct {
	ID        string   `json:"id"`
	Position  Position `json:"position"`
	Direction string   `json:"direction"` // "north", "east", "south", "west"
	Energy    int      `json:"energy"`
	Inventory []string `json:"inventory"`
	Actions   []Action `json:"actions"`
}

// MoveRequest is the payload for the move endpoint
type MoveRequest struct {
	Direction string `json:"direction"` // "up", "down", "left", "right"
}

// StateUpdateRequest is the payload for the state update endpoint
type StateUpdateRequest struct {
	Energy   *int      `json:"energy,omitempty"`
	Position *Position `json:"position,omitempty"`
}
