const passinput = document.getElementById("password")
passinput.addEventListener("change", (e) => {
  const id = "pass-err"
  removeErrorMessage(id)

  const value = e.currentTarget.value
  if (value.length < 8) {
    createErrorMessage("Password must be more than 8 characters", passinput, id)
  } else {
    removeErrorMessage(id)
  }
})

const confPassinput = document.getElementById("confirm-password")
confPassinput.addEventListener("change", (e) => {
  const id = "conf-error"
  removeErrorMessage(id)

  const confValue = e.currentTarget.value
  const passValue = passinput.value

  if (confValue !== passValue) {
    createErrorMessage("Passwords must match", confPassinput, id)
  } else {
    removeErrorMessage(id)
  }
})

function createErrorMessage(message, el, id) {
  const container = document.createElement("div")
  container.id = id 
  const msg = document.createElement("p")
  msg.innerText = message
  msg.classList.add("text-red-500", "text-sm")
  container.append(msg)
  el.insertAdjacentElement("afterend", container)
}

function removeErrorMessage(id) {
  const container = document.getElementById(id)
  
  if (container) {
    container.remove()
  }
}