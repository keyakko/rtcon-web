"use strict"
const host = "10.0.0.3"
const port = "5500"

export default () => {
  console.log = (msg) => {
    const xhr = new XMLHttpRequest()

    xhr.open('POST', 'http://' + host + ':' + port + '/')
    xhr.setRequestHeader("Content-Type", "application/json")
    xhr.send(JSON.stringify({
      type: "log",
      msg: String(msg)
    }))
  }
  console.info = (msg) => {
    const xhr = new XMLHttpRequest()

    xhr.open('POST', 'http://' + host + ':' + port + '/')
    xhr.setRequestHeader("Content-Type", "application/json")
    xhr.send(JSON.stringify({
      type: "info",
      msg: String(msg)
    }))
  }
  console.warn = (msg) => {
    const xhr = new XMLHttpRequest()

    xhr.open('POST', 'http://' + host + ':' + port + '/')
    xhr.setRequestHeader("Content-Type", "application/json")
    xhr.send(JSON.stringify({
      type: "warn",
      msg: String(msg)
    }))
  }
  console.error = (msg) => {
    const xhr = new XMLHttpRequest()

    xhr.open('POST', 'http://' + host + ':' + port + '/')
    xhr.setRequestHeader("Content-Type", "application/json")
    xhr.send(JSON.stringify({
      type: "err",
      msg: String(msg)
    }))
  }
  window.addEventListener("unhandledrejection", (e) => console.error(e))
  window.onerror = (e) => console.error(e)
}
