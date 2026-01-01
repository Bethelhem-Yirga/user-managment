import { useState } from "react";
import { createUser, getUser, updateUser, deleteUser } from "./api/user";
import "./App.css";

function App() {
  const [user, setUser] = useState({ id: "", name: "", email: "" });
  const [result, setResult] = useState(null);

  const handleCreate = async () => setResult(await createUser(user));
  const handleGet = async () => setResult(await getUser(user.id));
  const handleUpdate = async () => setResult(await updateUser(user));
  const handleDelete = async () => setResult(await deleteUser(user.id));

  return (
    <div className="container">
<h1 className="text-4xl font-bold text-blue-600 text-center">
  User CRUD
</h1>

      <div className="form">
        <input
          placeholder="ID"
          value={user.id}
          onChange={e => setUser({ ...user, id: e.target.value })}
        />
        <input
          placeholder="Name"
          value={user.name}
          onChange={e => setUser({ ...user, name: e.target.value })}
        />
        <input
          placeholder="Email"
          value={user.email}
          onChange={e => setUser({ ...user, email: e.target.value })}
        />
      </div>

      <div className="buttons">
        <button className="create" onClick={handleCreate}>Create</button>
        <button className="read" onClick={handleGet}>Read</button>
        <button className="update" onClick={handleUpdate}>Update</button>
        <button className="delete" onClick={handleDelete}>Delete</button>
      </div>

      <pre className="result">{JSON.stringify(result, null, 2)}</pre>
    </div>
  );
}

export default App;
