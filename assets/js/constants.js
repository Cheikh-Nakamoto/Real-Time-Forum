export const CONSTANTS = {
    API_URL: "http://localhost:" + "1112"  + "/api",
    NAMES: ["firstname", "lastname"],
    CONTAINERS: ["error-container", "home-container", "login-container", "register-container"],
    LINKS: ["register-go-login", "register-go-home", "login-go-register", 'login-go-home', "home-go-login", "home-go-register", "error-go-home"],
    SECTIONS_USERS: ['online-users', 'offline-users'],
    BUTTONS: ['online', 'offline']
}

Object.freeze(CONSTANTS)