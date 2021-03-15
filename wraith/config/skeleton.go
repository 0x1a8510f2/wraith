package config

type ConfigSkeleton struct {
	Process struct {
		RespectExitSignals bool
	}
	Comms struct {
		JWT struct {
			PublicKey string
		}
		Transmitter struct {
			URLGenerator         string
			Trigger              string
			TriggerCheckInterval int
			Encryption           struct {
				Enabled bool
				Type    int
				Key     string
			}
		}
		Receiver struct {
			URLGenerator         string
			Trigger              string
			TriggerCheckInterval int
			Encryption           struct {
				Enabled bool
				Type    int
				Key     string
			}
		}
	}
}
