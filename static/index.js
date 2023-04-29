// const participantBtn = document.querySelector('.participant');
const spectatorBtn = document.querySelector('.spectator');
const popup = document.querySelector('.popup');
const hidden = document.querySelector('.hidden');

spectatorBtn.addEventListener("click",  ()=>{
    popup.style.display = "none"
    hidden.style.display = "grid"
})