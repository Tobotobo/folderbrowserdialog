// folderbrowserdialog.go
// Copyright (c) 2021 Tobotobo
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package folderbrowserdialog

import (
	"github.com/lxn/win"

	"github.com/Tobotobo/shbrowseforfolder"
)

type FolderBrowserDialog struct {
	InnerValue folderBrowserDialog
}

type folderBrowserDialog struct {
	Owner               win.HWND
	Description         string
	RootFolder          string
	SelectedPath        string
	ShowNewFolderButton bool
}

func New() *FolderBrowserDialog {
	return &FolderBrowserDialog{
		InnerValue: folderBrowserDialog{
			Owner:               0,
			Description:         "フォルダーを指定してください。",
			RootFolder:          "",
			SelectedPath:        "",
			ShowNewFolderButton: true,
		},
	}
}

// ----------------------------------------------------------------

func (dlg *FolderBrowserDialog) Show() (accepted bool, selectedPath string) {
	owner := dlg.InnerValue.Owner
	title := dlg.InnerValue.Description
	rootDirPath := dlg.InnerValue.RootFolder
	initSelectedPath := dlg.InnerValue.SelectedPath
	flags := shbrowseforfolder.BIF_USENEWUI | shbrowseforfolder.BIF_VALIDATE

	if !dlg.InnerValue.ShowNewFolderButton {
		flags |= shbrowseforfolder.BIF_NONEWFOLDERBUTTON
	}

	accepted, selectedPath, err := shbrowseforfolder.Show(owner, title, rootDirPath, initSelectedPath, flags)
	if err != nil {
		panic(err)
	}
	dlg.InnerValue.SelectedPath = selectedPath

	return
}

func Show() (accepted bool, folderPath string) {
	dlg := New()
	return dlg.Show()
}

// ----------------------------------------------------------------

func (dlg *FolderBrowserDialog) Owner(owner win.HWND) *FolderBrowserDialog {
	dlg.InnerValue.Owner = owner
	return dlg
}

func Owner(owner win.HWND) *FolderBrowserDialog {
	dlg := New()
	return dlg.Owner(owner)
}

// ----------------------------------------------------------------

func (dlg *FolderBrowserDialog) Description(description string) *FolderBrowserDialog {
	dlg.InnerValue.Description = description
	return dlg
}

func Description(description string) *FolderBrowserDialog {
	dlg := New()
	return dlg.Description(description)
}

// ----------------------------------------------------------------

func (dlg *FolderBrowserDialog) RootFolder(path string) *FolderBrowserDialog {
	dlg.InnerValue.RootFolder = path
	return dlg
}

func RootFolder(path string) *FolderBrowserDialog {
	dlg := New()
	return dlg.RootFolder(path)
}

// ----------------------------------------------------------------

func (dlg *FolderBrowserDialog) SelectedPath(path string) *FolderBrowserDialog {
	dlg.InnerValue.SelectedPath = path
	return dlg
}

func SelectedPath(path string) *FolderBrowserDialog {
	dlg := New()
	return dlg.SelectedPath(path)
}

// ----------------------------------------------------------------

func (dlg *FolderBrowserDialog) HideNewFolderButton() *FolderBrowserDialog {
	dlg.InnerValue.ShowNewFolderButton = false
	return dlg
}

func HideNewFolderButton() *FolderBrowserDialog {
	dlg := New()
	return dlg.HideNewFolderButton()
}
