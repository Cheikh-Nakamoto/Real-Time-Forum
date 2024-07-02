import { RegisterAndLogin } from './compenant.js'
import { login, logout } from './fetch.js';

let connectbtn = document.getElementById('user-connect');
let changebtn = document.getElementById('change');
let userconnect = document.querySelector('.connect-username');
let logoutbtn = document.getElementById('logout');
connectbtn.addEventListener('click', (e) => {
    e.preventDefault();
    if ((userconnect.textContent).trim() == `Login/Register` && location.pathname != '/login') {
        window.history.pushState(null, null, '/login')
        changebtn.innerHTML = "";
        changebtn.innerHTML = RegisterAndLogin();
        initialize();
    }
});

document.addEventListener('click', function(e) {
    if (e.target && e.target.id === '.go-register') {
        setupFormSwitching();
    }
});
logoutbtn.addEventListener('click', (e) => {
    e.preventDefault();
    logout();
    window.history.pushState(null, null, '/')
    changebtn.innerHTML = "";
    document.getElementById("communaut").click();
})

function initialize() {
    const loginPath = '/login';
    console.log("ici c'est ", location.pathname);
    if (location.pathname == loginPath) {
        console.log('login');

        // Partie pour login
        let loginForm = document.getElementById('login-form');
        let loginButton = document.getElementById('login-button');

        // Attacher l'événement de soumission au formulaire
        loginButton.addEventListener('click', (e) => {
            e.preventDefault();
            handleLoginSubmit(loginForm);
            console.log('click login btn')
        });
        // Gérer l'affichage des formulaires
        setupFormSwitching()
        let registerbtn = document.getElementById('register-button');
        let registerForm = document.getElementById('register-form');
        registerbtn.addEventListener('click', (e) => {
            e.preventDefault();
            handleRegisterSubmit(registerForm);

        })
    }
}

async function handleLoginSubmit(form) {
    const loginButton = document.getElementById('login-button');
    loginButton.disabled = true; // Désactiver le bouton

    const formData = new FormData(form);
    const data = Object.fromEntries(formData.entries());
    try {
        const response = await login(data);
        if (response.token) {
            document.getElementById("communaut").click();
        }
    } catch (error) {
        console.error("Login failed:", error.message);
        // Afficher un message d'erreur avec un bouton de réessai
        const loginMessage = document.querySelector('.login-message');
        if (loginMessage) {
            loginMessage.classList.remove("hide");
            loginMessage.classList.add("is-invalid");
            loginMessage.textContent = "La connexion a échoué, veuillez réessayer.";
            // Ajouter un bouton de réessai dans le message ou à côté
            const retryButton = document.createElement('button');
            retryButton.textContent = "Réessayer";
            retryButton.onclick = () => handleLoginSubmit(form); // Réessayer la connexion
            loginMessage.appendChild(retryButton);
        }
    } finally {
        loginButton.disabled = false; // Réactiver le bouton après la tentative
    }
}


function handleRegisterSubmit(form) {
    // Obtenir une référence au bouton de soumission du formulaire d'inscription
    const registerButton = document.querySelector('#register-button');
    registerButton.disabled = true; // Désactiver le bouton

    const formData = new FormData(form);
    const data = Object.fromEntries(formData.entries());

    fetch('http://localhost:1112/api/register', {
        method: 'POST',
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
    })
    .then((res) => res.json())
    .then((data) => {
        console.log(data);
        // Réactiver le bouton
        registerButton.disabled = false;
        // Optionnel : Gérer la réussite de l'inscription, par exemple en affichant un message ou en redirigeant l'utilisateur
    })
    .catch((err) => {
        console.error("Registration failed:", err);
        // Afficher un message d'erreur avec un bouton de réessai
        const registerMessage = document.querySelector('.register-message');
        if (registerMessage) {
            registerMessage.classList.remove("hide");
            registerMessage.classList.add("is-invalid");
            registerMessage.textContent = "L'inscription a échoué, veuillez réessayer.";
            // Ajouter un bouton de réessai dans le message ou à côté
            const retryButton = document.createElement('button');
            retryButton.textContent = "Réessayer";
            retryButton.onclick = () => handleRegisterSubmit(form); // Réessayer l'inscription
            registerMessage.appendChild(retryButton);
        }
    })
    .finally(() => {
        // Réactiver le bouton après la tentative, indépendamment du résultat
        registerButton.disabled = false;
    });

    console.log("Ces données ont été envoyées", data);
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