import { useState } from "react";
import { createUser, getUser, updateUser, deleteUser } from "./api/user";

function App() {
  const [user, setUser] = useState({ id: "", name: "", email: "" });
  const [result, setResult] = useState(null);

  const handleCreate = async () => {
    const res = await createUser(user);
    setResult(res);
  };

  const handleGet = async () => {
    const res = await getUser(user.id);
    setResult(res);
  };

  const handleUpdate = async () => {
    const res = await updateUser(user);
    setResult(res);
  };

  const handleDelete = async () => {
    const res = await deleteUser(user.id);
    setResult(res);
  };

  return (
    <div className="App">
      <h1>User CRUD</h1>
      <input placeholder="ID" value={user.id} onChange={e => setUser({ ...user, id: e.target.value })} />
      <input placeholder="Name" value={user.name} onChange={e => setUser({ ...user, name: e.target.value })} />
      <input placeholder="Email" value={user.email} onChange={e => setUser({ ...user, email: e.target.value })} />
      <div>
        <button onClick={handleCreate}>Create</button>
        <button onClick={handleGet}>Read</button>
        <button onClick={handleUpdate}>Update</button>
        <button onClick={handleDelete}>Delete</button>
      </div>
      <pre>{JSON.stringify(result, null, 2)}</pre>
    </div>
  );
}

export default App;
