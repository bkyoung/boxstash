package cmd

import (
    "boxstash/internal/boxstash/repository"
    "boxstash/internal/boxstash/repository/shared/db"
    "boxstash/internal/boxstash/service"
    "boxstash/internal/handler/api"
    "bufio"
    "context"
    "fmt"
    "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "os"
    "strings"
)

// Config holds all application configuration options
type Config struct {
   DatabaseDriver     string
   DatabaseURI        string
   HostName           string
   LetsEncryptEnabled bool
   LetsEncryptEmail   string
   LogFile            string
   LogFormat          string
   LogLevel           string
   ServerPort         string
   ServerProto        string
   TLSCert            string
   TLSKey             string
}

func init() {

    viper.SetDefault("database-driver", "sqlite3")
    viper.SetDefault("database-uri", "./boxstash.db")
    viper.SetDefault("hostname", "localhost")
    viper.SetDefault("lets-encrypt-enabled", false)
    viper.SetDefault("log-file", "stdout")
    viper.SetDefault("log-format", "text")
    viper.SetDefault("log-level", "error")
    viper.SetDefault("server-port", ":8080")
    viper.SetDefault("server-proto", "http")
    serverCmd.PersistentFlags().StringP("database-driver", "",viper.GetString("database-driver"),
       "db driver")
    serverCmd.PersistentFlags().StringP("database-uri", "", viper.GetString("database-uri"),
        "db connection string")
    serverCmd.PersistentFlags().StringP("hostname", "",viper.GetString("hostname"),
        "server hostname")
    serverCmd.PersistentFlags().BoolP("lets-encrypt-enabled", "",
        viper.GetBool("lets-encrypt-enabled"),
        "Enable autocert through Let's Encrypt")
    serverCmd.PersistentFlags().StringP("lets-encrypt-email", "",
        viper.GetString("lets-encrypt-email"),
        "email to register Let's Encrypt cert (" +
            "only useful if --lets-encrypt-enabled is used)")
    serverCmd.PersistentFlags().StringP("log-file", "",viper.GetString("log-file"),
       "path to log file or stdout")
    serverCmd.PersistentFlags().StringP("log-format", "",viper.GetString("log-format"),
        "json/text")
    serverCmd.PersistentFlags().StringP("log-level", "",viper.GetString("log-level"),
        "log verbosity [error, info, debug]")
    serverCmd.PersistentFlags().StringP("server-port", "",viper.GetString("server-port"),
       "port to serve api on")
    serverCmd.PersistentFlags().StringP("server-proto", "", viper.GetString("server-proto"),
        "http/https")
    serverCmd.PersistentFlags().StringP("tls-cert","","","TLS certificate to use for HTTPS")
    serverCmd.PersistentFlags().StringP("tls-key","","","TLS private key for certificate")
    err := viper.BindPFlags(serverCmd.PersistentFlags())
    if err != nil {
        fmt.Errorf("ERROR %v", err)
    }
    viper.SetEnvPrefix("boxstash")
    viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
    viper.AutomaticEnv()
    rootCmd.AddCommand(serverCmd)
}

func loadConfig() (*Config, error) {
    return &Config{
        DatabaseDriver: viper.GetString("database-driver"),
        DatabaseURI: viper.GetString("database-uri"),
        HostName: viper.GetString("hostname"),
        LetsEncryptEnabled: viper.GetBool("lets-encrypt-enabled"),
        LetsEncryptEmail: viper.GetString("lets-encrypt-email"),
        LogFile: viper.GetString("log-file"),
        LogFormat: viper.GetString("log-format"),
        LogLevel: viper.GetString("log-level"),
        ServerPort: viper.GetString("server-port"),
        ServerProto: viper.GetString("server-proto"),
        TLSCert: viper.GetString("tls-cert"),
        TLSKey: viper.GetString("tls-key"),
    }, nil
}

func loadLogger(c *Config) *logrus.Logger {
    logger := logrus.New()

    // Configure log level
    level, err := logrus.ParseLevel(c.LogLevel)
    if err != nil {
        level = logrus.DebugLevel
    }
    logger.Level = level
    logger.Infof("log level set to %s", level)

    // Configure log output
    if c.LogFile != "" {
        if strings.ToLower(c.LogFile) == "stdout" {
            logger.Out = os.Stdout
            logger.Debug("logger writing to stdout")
        } else {
            f, errOpen := os.OpenFile(c.LogFile, os.O_RDWR|os.O_APPEND, 0660)
            if errOpen == nil {
                logger.Out = bufio.NewWriter(f)
                logger.Debugf("logger writing to ", c.LogFile)
            } else {
                logrus.StandardLogger().Errorf("Could not open logfile %s: %s", c.LogFile, errOpen)
                logger.Out = os.Stdout
                logger.Errorf("logger falling back to stdout")
            }
        }
    } else {
        logger.Info("logger writing to stdout")
    }

    // Configure log output format
    if strings.ToLower(c.LogFormat) == "json" {
        logger.Formatter = &logrus.JSONFormatter{
            DisableTimestamp: false,
        }
        logger.Debug("setting log format to json")
    } else {
        logger.Formatter = &logrus.TextFormatter{
            DisableTimestamp:          false,
            FullTimestamp:             true,
            DisableLevelTruncation:    true,
        }
        logger.Debug("setting log format to text")
    }

    return logger
}

func runServer() {
    cfg, err := loadConfig()
    if err != nil {
        fmt.Errorf("Error loading config: %v", err)
        os.Exit(1)
    }
    logger := loadLogger(cfg)
    db, err := db.Connect(cfg.DatabaseDriver, cfg.DatabaseURI)
    if err != nil {
        logger.WithField("func", "main").Panic("Error initializing database")
    }
    logger.WithFields(logrus.Fields{
        "func":   "main",
        "driver": cfg.DatabaseDriver,
        "datasource": cfg.DatabaseURI,
    }).Debug("Connected to database")

    BoxRepository := repository.NewBoxRepository(db, logger)
    BoxService := service.NewBoxService(BoxRepository, logger)
    BoxInteractor := api.NewInteractor(BoxService, logger)
    s := api.New(false, "jdoe@example.com", "", "", "localhost", ":8080",
        logger, BoxInteractor.Handler(), BoxInteractor)

    // TODO: Make a deadline and canceller
    ctx := context.Background()
    s.ListenAndServe(ctx)
}

var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Start boxstash in server mode",
    Long:  `Run the boxstash api server`,
    Run: func(cmd *cobra.Command, args []string) {
        runServer()
    },
}
