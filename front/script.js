document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const email = event.target.email.value;
    const password = event.target.password.value;

    // Backend API chaqiruvini amalga oshirish
    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            // Kirish muvaffaqiyatli bo'ldi, dashboard sahifasiga o'tish
            window.location.href = 'dashboard.html';
        } else {
            // Kirish muvaffaqiyatsiz bo'ldi, xatolikni ko'rsatish
            alert(data.message);
        }
    });
});

document.getElementById('signupForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const name = event.target.name.value;
    const surname = event.target.surname.value;
    const email = event.target.email.value;
    const password = event.target.password.value;
    const confirmPassword = event.target.confirm_password.value;

    if (password !== confirmPassword) {
        alert('Passwords do not match');
        return;
    }

    // Backend API chaqiruvini amalga oshirish
    fetch('/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name, surname, email, password })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            // Ro'yxatdan o'tish muvaffaqiyatli bo'ldi, login sahifasiga o'tish
            window.location.href = 'index.html';
        } else {
            // Ro'yxatdan o'tish muvaffaqiyatsiz bo'ldi, xatolikni ko'rsatish
            alert(data.message);
        }
    });
});
