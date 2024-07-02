import { checkAge, checkEmail, checkGender, checkName, checkNickname, checkPassword } from "./check.js";
import { CONSTANTS } from "./constants.js";
import { loadUsers, startSessionCheck } from "./fetch.js";

// document.body.addEventListener('click', (e) => {
//     if (location.pathname == '/login') {
//         console.log('login');

//         // partie pour login 
//         let loginform = document.getElementById('login-form');
//         let login = document.getElementById('login-button');
//         login.addEventListener('click', (e) => {
//             e.preventDefault();
//             const formdata = new FormData(loginform);
//             const data = Object.fromEntries(formdata.entries());
//             console.log("Ces donne on ete envoyer", data, formdata);
//             fetch('http://localhost:1112/api/login', {
//                 method: 'POST',
//                 headers: { "Content-Type": "application/json" },
//                 body: JSON.stringify(data),
//             }).then((res) => res.json())
//                 .then((data) => {
//                     console.log(data);
//                 })
//                 .catch((err) => {
//                     console.log(err);
//                 });
//         })

//         // Vérifiez si les écouteurs d'événements sont déjà attachés
//         if (!document.body.classList.contains('events-attached')) {
//             let registerButton = document.querySelector('.go-register');
//             registerButton.addEventListener('click', (e) => {
//                 console.log('click');
//                 document.getElementById('login-form').classList.add('hide');
//                 document.getElementById('register-form').classList.remove('hide');
//             });

//             let loginbtn = document.querySelector('.go-login');
//             loginbtn.addEventListener('click', (e) => {
//                 console.log('click');
//                 document.getElementById('login-form').classList.remove('hide');
//                 document.getElementById('register-form').classList.add('hide');
//             });

//             // Marquez que les écouteurs d'événements sont attachés
//             document.body.classList.add('events-attached');
//         }
//     } else if (location.pathname == '/register') {
//         let registerForm = document.getElementById('register-form');
//         registerForm.addEventListener('submit', (e) => {
//             e.preventDefault();
//             const formdata = new FormData(registerForm);
//             const data = Object.fromEntries(formdata.entries());
//             console.log("Ces donne on ete envoyer", data, formdata);
//             fetch('http://localhost:1112/api/register', {
//                 method: 'POST',
//                 headers: { "Content-Type": "application/json" },
//                 body: JSON.stringify(data),
//             }).then((res) => res.json())
//                 .then((data) => {
//                     console.log(data);
//                 })
//                 .catch((err) => {
//                     console.log(err);
//                 });
//             document.getElementById('login-form').classList.remove('hide');
//             document.getElementById('register-form').classList.add('hide');
//         });
//     }


// });

document.querySelector("#change").addEventListener('change', () => {
    console.log("Chargement complete !");
    initialize();
});

function initialize() {
    const loginPath = '/login';
    const registerPath = '/register';
    console.log("ici c'est ", location.pathname);
    if (location.pathname == loginPath) {
        console.log('login');

        // Partie pour login
        let loginForm = document.getElementById('login-form');
        let loginButton = document.getElementById('login-button');

        // Attacher l'événement de soumission au formulaire
        loginForm.addEventListener('submit', (e) => {
            e.preventDefault();
            handleLoginSubmit(loginForm);

        });

        // Gérer l'affichage des formulaires
        setupFormSwitching()
    } else if (location.pathname == registerPath) {
        checkregisterentries()
        console.log('register ma niggeu');
        let registerForm = document.getElementById('register-form');

        registerForm.addEventListener('submit', (e) => {
            e.preventDefault();
            handleRegisterSubmit(registerForm);
        });
    }
}

function handleLoginSubmit(form) {
    const loginButton = document.getElementById('login-button');
    loginButton.disabled = true; // Désactiver le bouton

    const formData = new FormData(form);
    const data = Object.fromEntries(formData.entries());
    console.log("Ces donnees ont ete envoyees", data);

    fetch('http://localhost:1112/api/login', {
        method: 'POST',
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    })
        .then((res) => res.json())
        .then((data) => {
            console.log(data);
            loginButton.disabled = false; // Réactiver le bouton
        })
        .catch((err) => {
            console.log(err);
            loginButton.disabled = false; // Réactiver le bouton en cas d'erreur
        });
}

function handleRegisterSubmit(form) {
    const registerButton = form.querySelector('button[type="submit"]');
    registerButton.disabled = true; // Désactiver le bouton

    const formData = new FormData(form);
    const data = Object.fromEntries(formData.entries());
    console.log("Ces donnees ont ete envoyees", data);

    fetch('http://localhost:1112/api/register', {
        method: 'POST',
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    })
        .then((res) => res.json())
        .then((data) => {
            console.log(data);
            registerButton.disabled = false; // Réactiver le bouton
        })
        .catch((err) => {
            console.log(err);
            registerButton.disabled = false; // Réactiver le bouton en cas d'erreur
        });

    document.getElementById('login-form').classList.remove('hide');
    document.getElementById('register-form').classList.add('hide');
}

function setupFormSwitching() {
    if (!document.body.classList.contains('events-attached')) {
        let registerButton = document.querySelector('.go-register');
        registerButton.addEventListener('click', () => {
            console.log('click');
            document.getElementById('login-form').classList.add('hide');
            document.getElementById('register-form').classList.remove('hide');
        });

        let loginButton = document.querySelector('.go-login');
        loginButton.addEventListener('click', () => {
            console.log('click');
            document.getElementById('login-form').classList.remove('hide');
            document.getElementById('register-form').classList.add('hide');
        });

        // Marquer que les écouteurs d'événements sont attachés
        document.body.classList.add('events-attached');
    }
}


// js de janel login et register ------------------------------------------------------------


export const checkRegister = (user) => {
    return checkNickname(user.nickname) && checkAge(`${user.age}`) && checkGender(user.gender) && checkEmail(user.email) && checkName(user.firstname) && checkName(user.lastname) && checkPassword(user.password)
}

export const checkLogin = (credentials) => {
    return (checkNickname(credentials.nickname) || checkEmail(credentials.nickname)) && checkPassword(credentials.password)
}

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

function checkregisterentries() {
    console.log("input ma nigeu ...")
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
}

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

    if (document.querySelector("form").getAttribute("method") === null || !/^POST$/gi.test(document.querySelector("form").getAttribute("method"))) {
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

