package tech

import "time"

type Segment struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
}

type UserSegment struct {
	Id       int       `json:"id" db:"id"`
	UserId   int       `json:"user_id" db:"user_id"`
	Segments []Segment `json:"segments" db:"segments" binding:"required"`
	AddedAt  time.Time `json:"added_at" db:"added_at"`
}

type USegments struct {
	Id          int    `json:"id" db:"id"`
	UserId      int    `json:"user_id" db:"user_id"`
	SegmentId   int    `json:"segment_id" db:"segment_id" binding:"required"`
	SegmentName string `json:"segment_name" bd:"segment_name" binding:"required"`
}
