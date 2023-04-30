// Get elements
const startStreamBtn = document.querySelector('.start-stream-btn');
const changeDataBtn = document.querySelector('.change-data-btn');
const changeDataForm = document.querySelector('.change-data-form');
const taskList = document.querySelector('#task-list');
const usernameSpan = document.querySelector('#username');
const phoneSpan = document.querySelector('#phone');
const emailSpan = document.querySelector('#email');
const newUsernameInput = document.querySelector('#new-username');
const newPhoneInput = document.querySelector('#new-phone');
const newEmailInput = document.querySelector('#new-email');
const newImageInput = document.querySelector('#new-image');
const userImage = document.querySelector('#userImage')

// Event listeners
startStreamBtn.addEventListener('click', () => {
    window.location.href = '/streaming';
});

changeDataBtn.addEventListener('click', () => {
    changeDataForm.classList.toggle('hidden');
});

changeDataForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('Username', newUsernameInput.value);
    formData.append('Phone', newPhoneInput.value);
    formData.append('Email', newEmailInput.value);
    formData.append('Image', newImageInput.files[0] ? newImageInput.files[0] : null);

    try {
        const response = await fetch('/api/user/changeData', {
            method: 'POST',
            body: formData
        });
        const data = await response.json();
        if (response.ok) {
            usernameSpan.textContent = data.Username;
            phoneSpan.textContent = data.Phone;
            emailSpan.textContent = data.Email;
            userImage.src = data.Image

            changeDataForm.classList.add('hidden');
        } else {
            console.log(data.error);
        }
    } catch (error) {
        console.log(error);
    }
});

// Functions
async function getTaskData() {
    try {
        const response = await fetch('/api/user/tasks');
        const data = await response.json();
        if (response.ok) {
            data.forEach((task) => {
                console.log(task)
                const taskItem = document.createElement('li');
                const taskName = document.createElement('h3');
                const taskDesc = document.createElement('p');
                taskName.textContent = task.Name;
                taskDesc.textContent = task.Description;
                taskItem.appendChild(taskName);
                taskItem.appendChild(taskDesc);
                taskList.appendChild(taskItem);
            });
        } else {
            console.log(data.error);
        }
    } catch (error) {
        console.log(error);
    }
}

async function getUserData() {
    try {
        const response = await fetch('/api/user/getUser');
        const data = await response.json();
        if (response.ok) {
            usernameSpan.textContent = data.Username;
            phoneSpan.textContent = data.Phone;
            emailSpan.textContent = data.Email;
            userImage.src = data.Image

            newUsernameInput.value = data.Username;
            newPhoneInput.value = data.Phone;
            newEmailInput.value = data.Email;

        } else {
            console.log(data.error);
        }
    } catch (error) {
        console.log(error);
    }
}

// Load initial data
getTaskData();
getUserData();
