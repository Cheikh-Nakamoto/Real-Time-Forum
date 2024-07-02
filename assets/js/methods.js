import {checkAge, checkEmail, checkGender, checkName, checkNickname, checkPassword} from "./check.js";
import {CONSTANTS} from "./constants.js";
import {loadUsers, startSessionCheck} from "./fetch.js";

export const checkRegister = (user) => {
    return checkNickname(user.nickname) && checkAge(`${user.age}`) && checkGender(user.gender) && checkEmail(user.email) && checkName(user.firstname) && checkName(user.lastname) && checkPassword(user.password)
}

export const checkLogin = (credentials) => {
    return (checkNickname(credentials.nickname) || checkEmail(credentials.nickname)) && checkPassword(credentials.password)
}

export const gotoPage = (tag) => {
    loadUsers()
    //startSessionCheck()
    CONSTANTS.CONTAINERS.forEach(container => {
        if (container === tag) {
            document.querySelector(`.${container}`).classList.remove("hide")
        } else {
            document.querySelector(`.${container}`).classList.add("hide")
        }
    })
}

export const onOrOfflineUsers = (tag) => {
    loadUsers()
    CONSTANTS.SECTIONS_USERS.forEach(section => {
        console.log(section, tag)
        if (section === tag) {
            document.querySelector(`#${section}`).classList.remove("hide")
        } else {
            document.querySelector(`#${section}`).classList.add("hide")
        }
    })
}

export const config = () => {
    const body = document.body
    body.onselectstart = () => false
    body.addEventListener("contextmenu", (e) => {
        e.preventDefault()
        gotoPage('home-container')
    })
}

export const userList = (obj) => {
    const li = document.createElement("li")
    const img = document.createElement("img")
    img.src = obj.image
    img.width = 40
    img.height = 40
    img.alt = `${obj.nickname}_${obj.id}`

    const div = document.createElement("div")

    const name = document.createElement('span')
    name.id = 'name'
    name.textContent = `${obj.firstname} ${obj.lastname.toUpperCase()}`

    const nick = document.createElement('span')
    nick.id = 'nick'
    nick.textContent = `@${obj.nickname}`
    div.append(name, nick)
    li.append(img, div)

    return li
}