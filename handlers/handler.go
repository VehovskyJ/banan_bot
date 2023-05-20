package handlers

import (
	"github.com/bwmarrin/discordgo"
	"padisoft/banana_farmer_bot/database"
)

type Handler struct {
	s  *discordgo.Session
	m  *discordgo.MessageCreate
	db *database.Database
}

func NewHandler(s *discordgo.Session, m *discordgo.MessageCreate, db *database.Database) *Handler {
	return &Handler{
		s:  s,
		m:  m,
		db: db,
	}
}
