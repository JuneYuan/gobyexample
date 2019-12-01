package main

import (
	"encoding/json"
	"time"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

// Simple solution: use an auxiliary struct with correctly formatted values
// Output: {"id":1,"name":"Ken","lastSeen":1575202738}
//
//func (u *MyUser) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&struct {
//		ID       int64     `json:"id"`
//		Name     string    `json:"name"`
//		LastSeen int64 `json:"lastSeen"`
//	}{
//		ID: u.ID,
//		Name: u.Name,
//		LastSeen: u.LastSeen.Unix(),
//	})
//}

// Technique: struct embed & alias
// Output: {"id":1,"name":"Ken","lastSeen":1575203098}
func (u *MyUser) MarshalJSON() ([]byte, error) {
	type Alias MyUser
	return json.Marshal(&struct {
		*Alias
		LastSeen int64 `json:"lastSeen"`
	}{
		Alias: (*Alias)(u),
		LastSeen: u.LastSeen.Unix(),
	})
}

func (u *MyUser) UnmarshalJSON(data []byte) error {
	type Alias MyUser
	aux := &struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.LastSeen = time.Unix(aux.LastSeen, 0)
	return nil
}