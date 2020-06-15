package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aoisensi/dota2-cog/models"
	"github.com/bwmarrin/discordgo"
)

func onMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	switch e.Message.Content {
	case "/dota2-cog debug delete-roles":
		commandDebugDeleteRoles(s, e)
	case "/dota2-cog debug hi":
		commandDebugHi(s, e)
	case "/dota2-cog register":
		commandRegister(s, e)
	case "/dota2-cog force-fetch":
		commandForceFetch(s, e)
	}
}

func commandRegister(s *discordgo.Session, e *discordgo.MessageCreate) {
	id, _ := strconv.ParseInt(e.Author.ID, 10, 63)
	nonce := makeConnectNonce(id)

	mes := fmt.Sprintf(
		"%v\n%v/connect?nonce=%v",
		"Please link your account from this URL.",
		env.Host, nonce,
	)
	s.ChannelMessageSend(e.ChannelID, mes)
}

func commandForceFetch(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.GuildID == "" {
		return
	}
	dGuild, err := s.Guild(e.GuildID)
	if err != nil {
		log.Println(err)
		return
	}
	admin := false
	for _, roleID := range e.Member.Roles {
		for _, role := range dGuild.Roles {
			if roleID == role.ID {
				if (role.Permissions & 0x8) > 0 {
					admin = true
					break
				}
			}
		}
		if admin {
			break
		}
	}
	if !admin {
		s.ChannelMessageSend(e.ChannelID, "This command is only available to the administrator.")
		return
	}
	if _, ok := fetching[e.GuildID]; ok {
		s.ChannelMessageSend(e.ChannelID, "I'm fetching. Don't spam it.")
		return
	}
	id, _ := strconv.ParseInt(e.GuildID, 10, 63)
	guild, err := models.FindGuild(context.Background(), db, id)
	if err != nil {
		log.Println(err)
		return
	}
	s.ChannelMessageSend(e.ChannelID, "Force fetching...")
	pid, _ := s.ChannelMessageSend(e.ChannelID, "▱▱▱▱▱▱▱▱▱▱ 0%")
	pc := make(chan progress)
	defer close(pc)
	go func() {
		for {
			p := <-pc
			x := p.Done * 10 / p.All
			msg := fmt.Sprintf(
				"%v%v %v%%",
				strings.Repeat("▰", x),
				strings.Repeat("▱", 10-x),
				p.Done*100/p.All,
			)
			s.ChannelMessageEdit(e.ChannelID, pid.ID, msg)
			if p.Done == p.All {
				break
			}
		}
	}()
	result, err := watchGuild(s, guild, pc)
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(e.ChannelID, "Failed! Sorry...")
		return
	}
	pc <- progress{Done: 1, All: 1} //force close
	s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Done!! Success: %v  Error: %v", result.Success, result.Error))

}

func commandDebugDeleteRoles(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Member.User.ID != "135617831864762368" {
		return
	}
	guild, err := s.Guild(e.GuildID)
	if err != nil {
		log.Println(err)
		return
	}
	for _, role := range guild.Roles {
		if strings.HasPrefix(role.Name, "Dota2 ") && role.Name != "Dota2 Cog" {
			if err := s.GuildRoleDelete(e.GuildID, role.ID); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

func commandDebugHi(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Member.User.ID != "135617831864762368" {
		return
	}
	s.ChannelMessageSend(e.ChannelID, "sup")
}
