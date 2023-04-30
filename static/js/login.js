const form = document.querySelector('form');
const username = document.querySelector('#username');
const password = document.querySelector('#password');

form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const url = '/api/user/login';
    const data = {
        username: username.value,
        password: password.value
    };

    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });

    if (response.ok) {
        window.location.href = 'profile';
    } else {
        alert('Incorrect username or password');
    }
});