package events

type EventType string

const (
	CHANNEL_CREATE      EventType = "CHANNEL_CREATE"
	CHANNEL_UPDATE      EventType = "CHANNEL_UPDATE"
	CHANNEL_DELETE      EventType = "CHANNEL_DELETE"
	GUILD_CREATE        EventType = "GUILD_CREATE"
	GUILD_UPDATE        EventType = "GUILD_UPDATE"
	GUILD_DELETE        EventType = "GUILD_DELETE"
	GUILD_EMOJIS_UPDATE EventType = "GUILD_EMOJIS_UPDATE"
	GUILD_MEMBER_ADD    EventType = "GUILD_MEMBER_ADD"
	GUILD_MEMBER_REMOVE EventType = "GUILD_MEMBER_REMOVE"
	GUILD_MEMBER_UPDATE EventType = "GUILD_MEMBER_UPDATE"
	GUILD_MEMBERS_CHUNK EventType = "GUILD_MEMBERS_CHUNK"
	GUILD_ROLE_CREATE   EventType = "GUILD_ROLE_CREATE"
	GUILD_ROLE_UPDATE   EventType = "GUILD_ROLE_UPDATE"
	GUILD_ROLE_DELETE   EventType = "GUILD_ROLE_DELETE"
	/*MESSAGE_CREATE             EventType = "MESSAGE_CREATE"
	MESSAGE_DELETE             EventType = "MESSAGE_DELETE"
	MESSAGE_DELETE_BULK        EventType = "MESSAGE_DELETE_BULK"
	MESSAGE_REACTION_ADD       EventType = "MESSAGE_REACTION_ADD"
	MESSAGE_REATION_REMOVE     EventType = "MESSAGE_REACTION_REMOVE"
	MESSAGE_REATION_REMOVE_ALL EventType = "MESSAGE_REACTION_REMOVE_ALL"*/
	VOICE_STATE_UPDATE  EventType = "VOICE_STATE_UPDATE"
	VOICE_SERVER_UPDATE EventType = "VOICE_SERVER_UPDATE"
)
