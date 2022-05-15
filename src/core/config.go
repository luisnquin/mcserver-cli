package core

import (
	"encoding/json"
	"os"

	"github.com/ProtonMail/go-appdir"
	jsoniter "github.com/json-iterator/go"
	"github.com/luisnquin/mcserver-cli/src/constants"
)

type Config struct {
	Apps   AppConfig `json:"-"`
	Config string    `json:"-"`
	Cache  string    `json:"-"`
	Data   string    `json:"-"`
	Logs   string    `json:"-"`
}

func LoadConfig(filePath string) *Config {
	app := appdir.New(appname)

	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var server AppConfig

	if err = jsoniter.Unmarshal(file, &server); err != nil {
		panic(err)
	}

	return &Config{
		Config: app.UserConfig(),
		Cache:  app.UserCache(),
		Data:   app.UserData(),
		Logs:   app.UserLogs(),
		Apps:   server,
	}
}

func OverwriteAppConfig(config *AppConfig) error {
	bytes, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(constants.ConfigFilePath, bytes, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

type AppConfig struct {
	Versions []string     `json:"versions"`
	Servers  []ServerData `json:"servers"`
}

type ServerData struct {
	Name    string       `json:"name"`
	Version string       `json:"version"`
	Config  ServerConfig `json:"config"`
}

type ServerConfig struct {
	EnableJmxMonitoring            bool   `json:"enable-jmx-monitoring" properties:"enable-jmx-monitoring,default=false"`
	RconPort                       int    `json:"rcon.port" properties:"rcon.port,default=25575"`
	LevelSeed                      string `json:"level-seed" properties:"level-seed"`
	Gamemode                       string `json:"gamemode,default=survival" properties:"gamemode,default=survival"`
	EnableCommandBlock             bool   `json:"enable-command-block,default" properties:"enable-command-block,default=false"`
	EnableQuery                    bool   `json:"enable-query,default" properties:"enable-query,default=false"`
	LevelName                      string `json:"level-name" properties:"level-name,default=world"`
	Modt                           string `json:"modt" properties:"modt,default=A Minecraft Server"`
	QueryPort                      int    `json:"query.port" properties:"query.port,default=25565"`
	Pvp                            bool   `json:"pvp" properties:"pvp,default=true"`
	GenerateStructures             bool   `json:"generate-structures" properties:"generate-structures,default=true"`
	Difficulty                     string `json:"difficulty" properties:"difficulty,default=easy"`
	NetworkCompressionThreshold    int    `json:"network-compression-threshold" properties:"network-compression-threshold,default=256"`
	MaxTickTime                    int    `json:"max-tick-time" properties:"max-tick-time,default=60000"`
	RequireResourcePack            bool   `json:"require-resource-pack" properties:"require-resource-pack,default=false"` // URI
	UseNativeTransport             bool   `json:"use-native-transport" properties:"use-native-transport,default=true"`
	MaxPlayers                     int    `json:"max-players" properties:"max-players,default=20"`
	OnlineMode                     bool   `json:"online-mode" properties:"online-mode,default=false"` // Allows no premium players
	EnableStatus                   bool   `json:"enable-status" properties:"enable-status,default=true"`
	AllowFlight                    bool   `json:"allow-flight" properties:"allow-flight,default=false"`
	BroadcatRconToOps              bool   `json:"broadcast-rcon-to-ops" properties:"broadcast-rcon-to-ops,default=true"`
	ViewDistance                   int    `json:"view-distance" properties:"view-distance,default=10"`
	ServerIp                       string `json:"server-ip" properties:"server-ip"`
	ResourePackPrompt              string `json:"resource-pack-prompt" properties:"resource-pack-prompt"`
	AllowNether                    bool   `json:"allow-nether" properties:"allow-nether,default=true"`
	ServerPort                     int    `json:"server-port" properties:"server-port,default=25565"`
	EnableRcon                     bool   `json:"enable-rcon" properties:"enable-rcon,default=false"`
	SyncChunkWrites                bool   `json:"sync-chunk-writes" properties:"sync-chunk-writes,default=true"`
	OpPermissionLevel              int    `json:"op-permission-level" properties:"op-permission-level,default=4"`
	PreventProxyConnections        bool   `json:"prevent-proxy-connections" properties:"prevent-proxy-connections,default=false"`
	HideOnlinePlayers              bool   `json:"hide-online-players" properties:"hide-online-players,default=false"`
	ResourcePack                   string `json:"resource-pack" properties:"resource-pack"`
	EntityBroadcastRangePercentage int    `json:"entity-broadcast-range-percentage" properties:"entity-broadcast-range-percentage,default=100"`
	SimulationDistance             int    `json:"simulation-distance" properties:"simulation-distance,default=10"`
	RconPassword                   string `json:"rcon.password" properties:"rcon.password"`
	PlayerIdleTimeout              int    `json:"player-idle-timeout" properties:"player-idle-timeout,default=0"`
	ForceGamemode                  bool   `json:"force-gamemode" properties:"force-gamemode,default=false"`
	RateLimit                      int    `json:"rate-limit" properties:"rate-limit,default=0"`
	Hardcode                       bool   `json:"hardcore" properties:"hardcore,default=false"`
	WhiteList                      bool   `json:"white-list" properties:"white-list,default=false"`
	BroadcastConsoleToOps          bool   `json:"broadcast-console-to-ops" properties:"broadcast-console-to-ops,default=true"`
	SpawnNpcs                      bool   `json:"spawn-npcs" properties:"spawn-npcs,default=true"`
	SpawnAnimals                   bool   `json:"spawn-animals" properties:"spawn-animals,default=true"`
	FunctionPermissionLevel        int    `json:"function-permission-level" properties:"function-permission-level,default=2"`
	LevelType                      string `json:"level-type" properties:"level-type,default=default"`
	TextFilteringConfig            string `json:"text-filtering-config" properties:"text-filtering-config"`
	SpawnMonsters                  bool   `json:"spawn-monsters" properties:"spawn-monsters,default=true"`
	EnforceWhitelist               bool   `json:"enforce-whitelist" properties:"enforce-whitelist,default=false"`
	ResourcePackSha1               string `json:"resource-pack-sha1" properties:"resource-pack-sha1"`
	SpawnProtection                int    `json:"spawn-protection" properties:"spawn-protection,default=16"`
	MaxWorldSize                   int    `json:"max-world-size" properties:"max-world-size,default=29999984"`
}
