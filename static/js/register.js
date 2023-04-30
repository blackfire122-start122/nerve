const registrationForm = document.getElementById('registration-form');

registrationForm.addEventListener('submit', (event) => {
    event.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const email = document.getElementById('email').value;
    const phone = document.getElementById('phone').value;

    // Відправляємо POST-запит на сервер
    fetch('/api/user/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username,
            password,
            email,
            phone
        })
    })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            window.location.href = 'login';
        })
        .catch(error => console.error(error));
});
