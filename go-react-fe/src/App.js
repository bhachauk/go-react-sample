import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import CreateUser from './components/CreateUser';
import EditUser from './components/EditUser';
import DeleteUser from './components/DeleteUser';
import UserTable from './components/UserTable';
import * as userService from './services/userService'; // Import user service
import './App.css'; // Import your CSS styles

function App() {
  const [users, setUsers] = useState([]);

  const fetchUsers = async () => {
    try {
      const response = await userService.getAllUsers();
      setUsers(response.data);
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  return (
      <Router>
        <div className="app-container">
          <nav className="tabs">
            <Link to="/create" className="tab-item">Create User</Link>
            <Link to="/edit" className="tab-item">Edit User</Link>
            <Link to="/delete" className="tab-item">Delete User</Link>
          </nav>

          <div className="content">
            <Routes>
              <Route path="/create" element={<CreateUser fetchUsers={fetchUsers} />} />
              <Route path="/edit" element={<EditUser users={users} fetchUsers={fetchUsers} />} />
              <Route path="/delete" element={<DeleteUser users={users} fetchUsers={fetchUsers} />} />
              <Route path="/" element={<UserTable users={users} />} /> {/* Default view */}
            </Routes>
          </div>
        </div>
      </Router>
  );
}

export default App;