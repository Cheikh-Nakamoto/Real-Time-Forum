export const checkNickname = (nickname = "") => {
    if (nickname.trim().length >= 3 && nickname.trim().length <= 20) {
        if (/^\d+/.test(nickname)) return false
        return /^[a-z0-9]+$/.test(nickname)
    }
    return false
}

export const checkAge = (value) => {
    if (value.trim().length === 0) return false
    const age = parseInt(value)
    if (isNaN(age)) return false
    return !(age < 12 || age > 120);

}

export const checkGender = (value) => {
    if (value.trim().length === 0) return false
    if (value.length < 4 || value.length > 6) return false
    return value === "male" || value === "female"
}

export const checkEmail = (value) => {
    if (/[A-Z]/.test(value)) return false
    if (value.length >= 10 && value.length <= 35) {
        const parts = value.split("@");
        
        if (parts.length !== 2 || parts[0].trim().length === 0 || parts[1].trim().length === 0) {
            return false;
        }

        const username = parts[0];
        const regUsername = /^[a-z0-9]+$/;
        if (username.length < 3 || username.length > 20 ||!regUsername.test(username)) {
            return false;
        }

        const domainParts = parts[1].split(".");
        return !(domainParts.length !== 2 || domainParts[0].trim().length < 3 || domainParts[0].trim().length > 10 || domainParts[1].trim().length < 2 || domainParts[1].trim().length > 3);
    }
    return false;
}

export const checkName = (value) => {
    if (value.length >= 3 && value.length <= 20) {
        return /^[a-zA-Z\s-]*$/.test(value);
    }
    return false;
}

export const checkPassword = (value) => {
    if (value.length >= 8 && value.length <= 20) {
        const hasUppercase = /[A-Z]/.test(value);
        const hasLowercase = /[a-z]/.test(value);
        const hasDigit = /\d/.test(value);
        const hasSpecialChar = /[\W_]/.test(value);
        return hasUppercase && hasLowercase && hasDigit && hasSpecialChar;
    }
    return false;
}