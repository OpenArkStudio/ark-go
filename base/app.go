package base

type ARKAppType uint8

const (
	ARK_APP_DEFAULT = iota // none
	ARK_APP_MASTER         // master //cluster level start
	ARK_APP_ROUTER         // router, world & cluster middle layer
	//ARK_APP_OSS,              // oss
	//ARK_APP_DIR,              // dir
	//ARK_APP_LOG,              // log
	//ARK_APP_CLUTER_RANK,      //cluster rank
	//ARK_APP_CLUSTER_MAIL,     //cluster mail
	//ARK_PROC_CLUSTER_PUB,     //cluster public
)

const (
	ARK_APP_CLUSTER_MAX = iota + 29 // max of cluster
	ARK_APP_WORLD                   // world // zone level start
	ARK_APP_GAME                    // game
	ARK_APP_LOGIN                   // login
	ARK_APP_PROXY                   // proxy
	ARK_APP_DB                      // db-proxy
	//ARK_APP_RANK,             // rank
	//ARK_APP_PUB,              // public
	//ARK_APP_CS_PROXY,         // cs_proxy, produce cross-server things
)

const (
	ARK_APP_WORLD_MAX = iota + 199 // max of world
	ARK_APP_REDIS                  // Redis server // others start
	ARK_APP_MYSQL                  // MySQL server
)

const (
	ARK_APP_MAX = 255 // max of all processes
)
