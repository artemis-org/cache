package objects

import (
	"fmt"
	"github.com/artemis-org/cache/utils"
	"github.com/fatih/structs"
	"strconv"
)

type Message struct {
	Id string
	ChannelId string
	GuildId string
	Author User
	Member Member
	Content string
	Timestamp string
	EditedTimestamp string
	Tts bool `json:"-"`
	MentionEveryone bool
	Mentions []MessageMentionedUser
	MentionsRoles []int64
	Attachments []Attachment `json:"-"` // We don't want to cache these
	Embeds []Embed `json:"-"` // We don't want to cache bot messages either
	Reactions []Reaction
	Nonce string
	Pinned bool
	WebhookId string
	Type int
	Activity MessageActivity `json:"-"`
	Application MessageApplication `json:"-"`
}

// Mentions is an array of users with partial member
type MessageMentionedUser struct {
	*User
	Member Member
}

func (m *Message) KeyName() string {
	return fmt.Sprintf("cache:message:%s", m.Id)
}

func (m *Message) Serialize() map[string]map[string]interface{} {
	fields := make(map[string]map[string]interface{})

	for k, v := range structs.Map(m) {
		if k == "Author" {
			fields = utils.Append(fields, m.Author.Serialize())
			fields[m.KeyName()][k] = m.Author.Id
		} else if k == "Member" {
			fields = utils.Append(fields, m.Member.Serialize(m.GuildId))
			// No need to cache member because we have guild ID + user ID
		} else if k == "Mentions" {
			// We only need to store user IDs, we have the rest of the data cached already
			var ids []string
			for _, user := range m.Mentions {
				ids = append(ids, user.Id)
			}
			fields[m.KeyName()][k] = utils.JoinToString(ids, ",")
		} else if k == "MentionsRoles" {
			var ids []string
			for _, id := range m.MentionsRoles {
				ids = append(ids, strconv.Itoa(int(id)))
			}
			fields[m.KeyName()][k] = utils.JoinToString(ids, ",")
		} else if k == "Reactions" {
			var ids []string
			for _, reaction := range m.Reactions {
				fields = utils.Append(fields, reaction.Serialize(m.Id))
				ids = append(ids, reaction.Emoji.Id)
			}
			fields[m.KeyName()][k] = utils.JoinToString(ids, ",")
		} else {
			fields[m.KeyName()][k] = v
		}
	}

	return fields
}
