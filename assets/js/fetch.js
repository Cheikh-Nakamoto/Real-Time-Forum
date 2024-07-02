import { CONSTANTS } from "./constants.js"
import { gotoPage, userList } from "./methods.js";
import { FetchAllPosts,DispatchePost } from "./home.js";
import {Animation} from "./animation.js"

export const register = async (user) => {
    const response = await fetch(`${CONSTANTS.API_URL}/register`, {
        method: "POST",
        mode: "cors",
        headers: {
            "Accept": "multipart/form-data",
            "Content-Type": "application/json",
        },
        body: JSON.stringify(user),
        credentials: "include",
    })
}

// Async login function
export const login = async (credentials) => {
    const response = await fetch(`${CONSTANTS.API_URL}/login`, {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(credentials),
    })

    const data = await response.json()
    if (data.token) {
        localStorage.setItem("token", data.token)
        localStorage.setItem("object", data.object)
    }
    UserInf()
    return data
}

// Async logout function
export const logout = async () => {
    const token = localStorage.getItem("object")
    const response = await fetch(`${CONSTANTS.API_URL}/session`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ "ID": token }),
    })
    if (response.status === 200) {
        localStorage.removeItem("token")
        localStorage.removeItem("object")
        document.querySelector(".connect-username").innerHTML = "<h3>Login/Register</h3>"
        document.getElementById("communaut").click();
        FetchAllPosts().then(data => {
            DispatchePost(data)
            Animation();
        })
    }
    return response.status === 200
}

// Async get online users function
export const getOnlineUsers = async () => {
    return await fetch(`${CONSTANTS.API_URL}/online-users`, {
        method: "GET",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
        }
    }).then(response => {
        if (response.ok) {
            return response.json()
        }
        throw new Error("Failed to get online users.")
    }).then(data => {
        return data
    })
        .catch(error => {
            throw error
        })

}

// Async get offline users function
export const getOfflineUsers = async () => {
    return await fetch(`${CONSTANTS.API_URL}/offline-users`, {
        method: "GET",
        headers: {
            "Accept": "application/json",
            "Content-Type": "application/json",
        }
    }).then(response => {
        if (response.ok) {
            return response.json()
        }
        throw new Error("Failed to get offline users.")
    }).then(data => {
        return data
    })
        .catch(error => {
            throw error
        })
}

export const loadUsers = async () => {
    try {
        await getOnlineUsers()
            .then(data => {
                if (Array.isArray(data)) {
                    const onlineUsersList = document.getElementById("online-users")
                    onlineUsersList.innerHTML = ""
                    data.forEach(user => {
                        onlineUsersList.append(userList(user))
                    })
                } else {
                    const onlineUsersList = document.getElementById("online-users")
                    const li = document.createElement("li")
                    li.textContent = data.message
                    onlineUsersList.append(li)
                }
            }).catch(error => {
                console.error(error.message)
                return []
            })

        await getOfflineUsers()
            .then(data => {
                if (Array.isArray(data)) {
                    const offlineUsersList = document.getElementById("offline-users")
                    offlineUsersList.innerHTML = ""
                    data.forEach(user => {
                        offlineUsersList.append(userList(user))
                    })
                } else {
                    const offlineUsersList = document.getElementById("offline-users")
                    const li = document.createElement("li")
                    li.textContent = data.message
                    offlineUsersList.append(li)
                }
            }).catch(error => {
                console.error(error.message)
                return []
            })
    } catch (error) {
        console.error(error.message.toString())
    }
}

export const checkSession = async () => {
    const token = localStorage.getItem('token')
    if (!token) {
        //alert('You are not logged in.')
        gotoPage('login-container')
        return
    }

    const response = await fetch(`${CONSTANTS.API_URL}/check-session`, {
        method: 'GET',
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
    })
    if (response.status !== 200) {
        //alert('You are not logged in.')
        gotoPage('login-container')
    }
}

export const startSessionCheck = async () => {
    await checkSession()
    setInterval(checkSession, 5 * 60 * 1000)
}

export const UserInf = () => {
    let userID = localStorage.getItem('object')
    console.log(userID)
    fetch(`http://localhost:1112/api/user`, {
        method: 'POST',
        headers: {
            "Content-Type": "application/json",
        },
    
        body: JSON.stringify({ 'ID': userID }),
    }).then((res) => res.json()).then((data) => {
        document.querySelector(".connect-username").innerHTML = `<h3>${data.nickname}<h3/>`;
    })
}
