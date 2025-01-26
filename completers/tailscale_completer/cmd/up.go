package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Connect to Tailscale, logging in if needed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	upCmd.Flags().Bool("accept-dns", true, "accept DNS configuration from the admin panel (default true)")
	upCmd.Flags().String("accept-risk", "", "accept risk and skip confirmation for risk types: lose-ssh,mac-app-connector,all")
	upCmd.Flags().Bool("accept-routes", false, "accept routes advertised by other Tailscale nodes (default false)")
	upCmd.Flags().Bool("advertise-connector", false, "advertise this node as an app connector (default false)")
	upCmd.Flags().Bool("advertise-exit-node", false, "offer to be an exit node for internet traffic for the tailnet (default false)")
	upCmd.Flags().String("advertise-routes", "", "routes to advertise to other nodes (comma-separated, e.g. \"10.0.0.0/8,192.168.0.0/24\") or empty string to not advertise routes")
	upCmd.Flags().String("advertise-tags", "", "comma-separated ACL tags to request; each must start with \"tag:\" (e.g. \"tag:eng,tag:montreal,tag:ssh\")")
	upCmd.Flags().String("auth-key", "", "node authorization key; if it begins with \"file:\", then it's a path to a file containing the authkey")
	upCmd.Flags().String("exit-node", "", "Tailscale exit node (IP or base name) for internet traffic, or empty string to not use an exit node")
	upCmd.Flags().Bool("exit-node-allow-lan-access", false, "Allow direct access to the local network when routing traffic via an exit node (default false)")
	upCmd.Flags().Bool("force-reauth", false, "force reauthentication (default false)")
	upCmd.Flags().String("hostname", "", "hostname to use instead of the one provided by the OS")
	upCmd.Flags().Bool("json", false, "output in JSON format (WARNING: format subject to change) (default false)")
	upCmd.Flags().String("login-server", "", "base URL of control server (default https://controlplane.tailscale.com)")
	upCmd.Flags().String("netfilter-mode", "", "netfilter mode (one of on, nodivert, off) (default on)")
	upCmd.Flags().String("operator", "", "Unix username to allow to operate on tailscaled without sudo")
	upCmd.Flags().Bool("qr", false, "show QR code for login URLs (default false)")
	upCmd.Flags().Bool("reset", false, "reset unspecified settings to their default values (default false)")
	upCmd.Flags().Bool("shields-up", false, "don't allow incoming connections (default false)")
	upCmd.Flags().Bool("snat-subnet-routes", true, "source NAT traffic to local routes advertised with --advertise-routes (default true)")
	upCmd.Flags().Bool("ssh", false, "run an SSH server, permitting access per tailnet admin's declared policy (default false)")
	upCmd.Flags().Bool("stateful-filtering", false, "apply stateful filtering to forwarded packets (subnet routers, exit nodes, etc.) (default false)")
	upCmd.Flags().String("timeout", "", "maximum amount of time to wait for tailscaled to enter a Running state; default (0s) blocks forever (default 0s)")
	rootCmd.AddCommand(upCmd)

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"accept-risk": carapace.ActionValues("lose-ssh", "mac-app-connector", "all").UniqueList(","),
		"netfilter-mode": carapace.ActionValues("on", "nodivert", "off"),
	})
}
