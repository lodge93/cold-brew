const {app, BrowserWindow} = require('electron')
const path = require('path')
const url = require('url')

function createWindow () {
  win = new BrowserWindow({width: 800, height: 600})
  win.loadURL("http://localhost:3000")
  win.webContents.openDevTools()
}

app.on('ready', createWindow)