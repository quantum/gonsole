package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
	g "github.com/quantum/gonsole"
)

func openProgressBar(app *g.App) {
	title := "Castle Installer - Installing CastleOS"
	message := "Please wait while the installation completes."

	w := g.NewWindow(app, "progress")
	w.SetPosition(g.Position{"15%", "25%", "70%", "40%"})
	w.SetTitle(title)

	label := g.NewLabel(w, w, fmt.Sprintf("%s__message", w.ID()))
	label.SetPosition(g.Position{"0", "0", "100%", "1"})
	label.SetText(message)

	prg := g.NewProgress(w, w, fmt.Sprintf("%s__progress", w.ID()))
	prg.SetPosition(g.Position{"0%", "50%-1", "100%", "1"})
	prg.SetValue(1.20)
}

func openJoinClusterDialog(app *g.App) {
	title := "Castle Installer - Join Cluster"
	message := "Enter the secret key to join the cluster."

	d := g.NewInputDialog(app, "join", title, message, []string{"Continue"})
	d.SetPosition(g.Position{"15%", "25%", "70%", "40%"})
	d.AddEventListener("closed", func(ev *g.Event) bool {
		openProgressBar(app)
		return true
	})
}

func openSystemDiskDialog(app *g.App) {
	title := "Castle Installer - Choose System Disk"
	message := "Select system disk"

	d := g.NewSelectionDialog(app, "system", title, message, []string{"Continue"}, []string{"sda", "sdb", "sdc", "sde", "sdf", "sdg", "sdh"})
	d.SetPosition(g.Position{"15%", "25%", "70%", "40%"})
	d.AddEventListener("closed", func(ev *g.Event) bool {
		openJoinClusterDialog(app)
		return true
	})
}

func openWelcomeDialog(app *g.App) {
	title := "Castle Installer - Welcome"
	message := "Welcome to the CastleOS installer. This installer will walk you through all " +
		"necessary steps needed to get this node connected to the cluster you created in " +
		"the web app. Once CastleOS is successfully installed and connected to the cluster, " +
		"you will have the ability to view and manage it via the web app. You will need to have " +
		"the clusters secret key on hand in order to join this node to the cluster. You will also " +
		"need to ensure this node is connected to the Internet so it can communicate and connect with " +
		"your Castle account. It is also important to note that the CastleOS will need to wipe all " +
		"disks on this node and if there is existing data it will be overwritten."

	mb := g.NewMessageDialog(app, "welcome", title, message, []string{"Start Installation"})
	mb.SetPosition(g.Position{"15%", "25%", "70%", "40%"})

	mb.AddEventListener("closed", func(ev *g.Event) bool {
		openSystemDiskDialog(app)
		return true
	})
}

func main() {
	app := g.NewApp()
	app.CloseKey = termbox.KeyEsc

	infoBar := g.NewWindow(app, "info")
	infoBar.SetBorderType(g.LineNone)
	infoBar.SetPosition(g.Position{"0", "100%-1", "100%", "2"})

	infoBarLabel := g.NewLabel(infoBar, infoBar, "infobar")
	infoBarLabel.SetText("[Tab] to move, [Enter] to select, [F1] for help, [F10] to exit")
	infoBarLabel.SetPosition(g.Position{"0", "0", "100%", "100%"})

	openProgressBar(app)

	app.Run()
}
