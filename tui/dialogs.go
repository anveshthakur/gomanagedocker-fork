package tui

import (
	"regexp"

	"github.com/ajayd-san/gomanagedocker/dockercmd"
	"github.com/ajayd-san/gomanagedocker/tui/components"
	teadialog "github.com/ajayd-san/teaDialog"
)

const (
	dialogRemoveContainer teadialog.DialogType = iota
	dialogPruneContainers
	dialogRemoveImage
	dialogPruneImages
	dialogRunImage
	dialogPruneVolumes
	dialogRemoveVolumes
	dialogImageScout
	dialogImageBuild
	dialogImageBuildProgress
)

func getRunImageDialog(storage map[string]string) teadialog.Dialog {
	prompt := []teadialog.Prompt{
		teadialog.MakeTextInputPrompt("port", "Port mappings"),
		teadialog.MakeTextInputPrompt("name", "Name"),
		teadialog.MakeTextInputPrompt("env", "Environment variables"),
	}

	return teadialog.InitDialogWithPrompt("Run Image", prompt, dialogRunImage, storage)
}

func getImageScoutDialog(f func() (*dockercmd.ScoutData, error)) InfoCardWrapperModel {
	infoCard := teadialog.InitInfoCard(
		"Image Scout",
		"",
		dialogImageScout,
		teadialog.WithMinHeight(13),
		teadialog.WithMinWidth(130),
	)
	return InfoCardWrapperModel{
		tableChan: make(chan *TableModel),
		inner:     &infoCard,
		f:         f,
		spinner:   components.InitialModel(),
	}
}

func getRemoveContainerDialog(storage map[string]string) teadialog.Dialog {
	prompts := []teadialog.Prompt{
		teadialog.MakeTogglePrompt("remVols", "Remove volumes?"),
		teadialog.MakeTogglePrompt("remLinks", "Remove links?"),
		teadialog.MakeTogglePrompt("force", "Force?"),
	}

	return teadialog.InitDialogWithPrompt("Remove Container Options:", prompts, dialogRemoveContainer, storage)
}

func getRemoveVolumeDialog(storage map[string]string) teadialog.Dialog {
	prompts := []teadialog.Prompt{
		teadialog.MakeTogglePrompt("force", "Force?"),
	}

	return teadialog.InitDialogWithPrompt("Remove Volume Options:", prompts, dialogRemoveVolumes, storage)
}

func getPruneContainersDialog(storage map[string]string) teadialog.Dialog {
	prompts := []teadialog.Prompt{
		teadialog.MakeOptionPrompt("confirm", "This will remove all stopped containers, are your sure?", []string{"Yes", "No"}),
	}

	return teadialog.InitDialogWithPrompt("Prune Containers: ", prompts, dialogPruneContainers, storage)
}

func getRemoveImageDialog(storage map[string]string) teadialog.Dialog {
	prompts := []teadialog.Prompt{
		teadialog.MakeTogglePrompt("force", "Force"),
		teadialog.MakeTogglePrompt("pruneChildren", "Prune Children"),
	}

	return teadialog.InitDialogWithPrompt("Remove Image Options:", prompts, dialogRemoveImage, storage)
}

func getPruneImagesDialog(storage map[string]string) teadialog.Dialog {
	prompts := []teadialog.Prompt{
		teadialog.MakeOptionPrompt("confirm", "This will remove all unused images, are your sure?", []string{"Yes", "No"}),
	}

	return teadialog.InitDialogWithPrompt("Prune Containers: ", prompts, dialogPruneImages, storage)
}

func getPruneVolumesDialog(storage map[string]string) teadialog.Dialog {
	prompts := []teadialog.Prompt{
		teadialog.MakeTogglePrompt("all", "Removed all unused volumes(not just anonymous ones)"),
		teadialog.MakeOptionPrompt("confirm", "This will remove all unused volumes, are your sure?", []string{"Yes", "No"}),
	}

	return teadialog.InitDialogWithPrompt("Prune Containers: ", prompts, dialogPruneVolumes, storage)
}

func getBuildImageDialog(storage map[string]string) teadialog.Dialog {
	prompts := []teadialog.Prompt{
		// teadialog.NewFilePicker("browser"),
		// NewFilePicker("filepicker"),
		teadialog.MakeTextInputPrompt("image_tags", "Image Tags:"),
	}

	return teadialog.InitDialogWithPrompt("Build Image: ", prompts, dialogImageBuild, storage)
}

// Gets the build progress bar info card/dialog
func getBuildProgress(progressBar components.ProgressBar) buildProgressModel {

	infoCard := teadialog.InitInfoCard(
		"Image Build",
		"",
		dialogImageBuildProgress,
		teadialog.WithMinHeight(8),
		teadialog.WithMinWidth(100),
	)

	reg := regexp.MustCompile(`Step\s(\d+)\/(\d+)\s:\s(.*)`)

	return buildProgressModel{
		progressChan: make(chan string, 10),
		regex:        reg,
		progressBar:  progressBar,
		inner:        &infoCard,
	}
}
