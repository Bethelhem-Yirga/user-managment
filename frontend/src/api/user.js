const BASE_URL = "http://localhost:8080";

export async function createUser(user) {
  const res = await fetch(`${BASE_URL}/users`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(user),
  });
  return res.json();
}

export async function getUser(id) {
  const res = await fetch(`${BASE_URL}/users/${id}`);
  return res.json();
}

export async function updateUser(user) {
  const res = await fetch(`${BASE_URL}/users/${user.id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(user),
  });
  return res.json();
}

export async function deleteUser(id) {
  const res = await fetch(`${BASE_URL}/users/${id}`, { method: "DELETE" });
  return res.json();
}
