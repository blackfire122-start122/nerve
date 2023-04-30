// const participantBtn = document.querySelector('.participant');
const spectatorBtn = document.querySelector('.spectator');
const popup = document.querySelector('.popup');
const hidden = document.querySelector('.hidden');
let user

function hidePopup(){
    popup.style.display = "none"
    hidden.style.display = "grid"
}

function showPopup(){
    popup.style.display = "block"
    hidden.style.display = "none"
}

spectatorBtn.addEventListener("click",  hidePopup)

async function getUserData() {
    try {
        const response = await fetch('/api/user/getUser');
        const data = await response.json();
        if (response.ok) {
            user = data
            hidePopup()
        } else {
            console.log(data.error);
        }
    } catch (error) {showPopup()}
}

getUserData()
