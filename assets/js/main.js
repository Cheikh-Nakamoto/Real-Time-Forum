import {getOfflineUsers, loadUsers, login, logout, register, startSessionCheck} from "./fetch.js"
import {
    checkAge,
    checkEmail,
    checkGender,
    checkName,
    checkNickname,
    checkPassword,
} from "./check.js"
import {config, checkRegister, checkLogin, gotoPage, onOrOfflineUsers} from "./methods.js"
import {CONSTANTS} from "./constants.js";

config()
await loadUsers()
//startSessionCheck()

CONSTANTS.NAMES.forEach((name) => {
    document.getElementById(name).addEventListener("input", (event) => {
        event.preventDefault()

        if (checkName(event.target.value)) {
            document.getElementById(name).classList.remove("is-invalid")
            document.getElementById(name).classList.add("is-valid")
        } else {
            document.getElementById(name).classList.remove("is-valid")
            document.getElementById(name).classList.add("is-invalid")
        }
    })
})

CONSTANTS.LINKS.forEach(link => {
    document.getElementById(link).addEventListener("click", (event) => {
        event.preventDefault()
        switch (link) {
            case `register-go-login`:
            case `home-go-login`:
                gotoPage("login-container")
                break
            case `login-go-register`:
            case `home-go-register`:
                gotoPage("register-container")
                break
            default:
                gotoPage("home-container")
                break
        }
    })
})

CONSTANTS.BUTTONS.forEach(button => {
    document.getElementById(button).addEventListener("click", (event) => {
        event.preventDefault()

        if (button === 'online') {
            document.querySelector('#offline-users').style.display = "none"
        } else {
            console.log('no')
        }
    })
})

document.getElementById("nickname").addEventListener("input", (event) => {
    event.preventDefault()

    if (checkNickname(event.target.value)) {
        document.getElementById("nickname").classList.remove("is-invalid")
        document.getElementById("nickname").classList.add("is-valid")
    } else {
        document.getElementById("nickname").classList.remove("is-valid")
        document.getElementById("nickname").classList.add("is-invalid")
    }
})

document.getElementById("age").addEventListener("input", (event) => {
    event.preventDefault()

    if (checkAge(event.target.value)) {
        document.getElementById("age").classList.remove("is-invalid")
        document.getElementById("age").classList.add("is-valid")
    } else {
        document.getElementById("age").classList.remove("is-valid")
        document.getElementById("age").classList.add("is-invalid")
    }
})

document.getElementById("gender").addEventListener("change", (event) => {
    event.preventDefault()

    if (event.target.value.trim().length !== 0) {
        if (checkGender(event.target.value)) {
            document.getElementById("gender").classList.remove("is-invalid")
            document.getElementById("gender").classList.add("is-valid")
        } else {
            document.getElementById("gender").classList.remove("is-valid")
            document.getElementById("gender").classList.add("is-invalid")
        }
    } else {
        document.getElementById("gender").classList.remove("is-valid")
        document.getElementById("gender").classList.add("is-invalid")
    }
})

document.getElementById("email").addEventListener("input", (event) => {
    event.preventDefault()

    if (checkEmail(event.target.value)) {
        document.getElementById("email").classList.remove("is-invalid")
        document.getElementById("email").classList.add("is-valid")
    } else {
        document.getElementById("email").classList.remove("is-valid")
        document.getElementById("email").classList.add("is-invalid")
    }
})

document.getElementById("password").addEventListener("input", (event) => {
    event.preventDefault()

    if (checkPassword(event.target.value)) {
        document.getElementById("password").classList.remove("is-invalid")
        document.getElementById("password").classList.add("is-valid")
    } else {
        document.getElementById("password").classList.remove("is-valid")
        document.getElementById("password").classList.add("is-invalid")
    }
})

document.getElementById("register-button").addEventListener("click", async (event) => {
    event.preventDefault()

    const user = {
        nickname: document.getElementById("nickname").value,
        age: parseInt(document.getElementById("age").value),
        gender: document.getElementById("gender").value,
        firstname: document.getElementById("firstname").value,
        lastname: document.getElementById("lastname").value,
        email: document.getElementById("email").value,
        password: document.getElementById("password").value,
    }

    if (!checkRegister(user)) {
        document.querySelector('.register-message').classList.remove("hide")
        document.querySelector('.register-message').classList.add("is-invalid")
        document.querySelector('.register-text-message').textContent = "Please fill in all fields correctly."

        setTimeout(() => {
            document.querySelector('.register-message').classList.remove("is-invalid")
            document.querySelector('.register-message').classList.add("hide")
        }, 5000)
        return
    }

    if (document.querySelector("form").getAttribute("method") === null || !/^POST$/gi.test(document.querySelector("form").getAttribute("method"))) {
        console.log("Registration failed. Please try again.")
        return
    }

    try {
        await register(user)
        document.querySelector('.register-container').classList.add("hide")
        document.querySelector('.login-container').classList.remove("hide")
    } catch (error) {
        console.error("Registration failed:", error.message)
    }
})

document.getElementById("username").addEventListener('input', (event) => {
    event.preventDefault()

    if (checkNickname(event.target.value) || checkEmail(event.target.value)) {
        document.getElementById("username").classList.remove("is-invalid")
        document.getElementById("username").classList.add("is-valid")
    } else {
        document.getElementById("username").classList.remove("is-valid")
        document.getElementById("username").classList.add("is-invalid")
    }
})

document.getElementById("pass").addEventListener('input', (event) => {
    event.preventDefault()

    if (checkPassword(event.target.value)) {
        document.getElementById("pass").classList.remove("is-invalid")
        document.getElementById("pass").classList.add("is-valid")
    } else {
        document.getElementById("pass").classList.remove("is-valid")
        document.getElementById("pass").classList.add("is-invalid")
    }
})

document.getElementById("login-button").addEventListener("click", async (event) => {
    event.preventDefault()

    const credentials = {
        nickname: document.getElementById("username").value,
        password: document.getElementById("pass").value,
    }

    if (credentials.nickname.trim().length === 0 || credentials.password.trim().length === 0) {
        document.querySelector('.login-message').classList.remove("hide")
        document.querySelector('.login-message').classList.add("is-invalid")
        document.querySelector('.login-text-message').textContent = "Fields cannot be empty."
        return
    }

    if (!checkLogin(credentials)) {
        document.querySelector('.login-message').classList.remove("hide")
        document.querySelector('.login-message').classList.add("is-invalid")
        document.querySelector('.login-text-message').textContent = "Please fill in all fields correctly."
        return
    }

    if (document.querySelector("form").getAttribute("method") === null ||!/^POST$/gi.test(document.querySelector("form").getAttribute("method"))) {
        console.log("Login failed. Please try again.")
        return
    }

    try {
        const response = await login(credentials);
        if (response.token) {
            gotoPage('home-container')
        }
    } catch (error) {
        console.error("Login failed:", error.message)
        document.querySelector('.login-message').classList.remove("hide")
        document.querySelector('.login-message').classList.add("is-invalid")
        document.querySelector('.login-text-message').textContent = "Incorrect nickname or password"
    }
})

document.getElementById("logout-button").addEventListener("click", async (event) => {
    event.preventDefault()

    const response = await logout()
    if (response) {
        gotoPage('home-container')
        console.log("Logging out...")
    } else {
        console.log("Logout failed.")
    }
})
