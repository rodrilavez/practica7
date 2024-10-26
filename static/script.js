const userList = document.getElementById("user-list");
const userForm = document.getElementById("user-form");

userForm.addEventListener("submit", async (event) => {
  event.preventDefault();
  const formData = new FormData(userForm);
  await fetch('/api/users', {
    method: 'POST',
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(Object.fromEntries(formData))
  });
  getUsers();
});

async function getUsers() {
  const response = await fetch('/api/users');
  const users = await response.json();
  userList.innerHTML = '';
  users.forEach(user => {
    const li = document.createElement('li');
    const btn = document.createElement('button');
    btn.textContent = 'Eliminar';
    btn.addEventListener('click', () => deleteUser(user.id));
    li.textContent = `${user.name} - ${user.email}`;
    li.appendChild(btn);
    userList.appendChild(li);
  });
}

async function deleteUser(id) {
  await fetch(`/api/users/${id}`, { method: 'DELETE' });
  getUsers();
}

getUsers();
