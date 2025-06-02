import React, { useState } from 'react';
import * as userService from '../services/userService';

function DeleteUser({ users, fetchUsers }) {
    const [selectedUserId, setSelectedUserId] = useState('');
    const [message, setMessage] = useState('');

    const handleSelectChange = (e) => {
        setSelectedUserId(e.target.value);
        setMessage('');
    };

    const handleDelete = async (e) => {
        e.preventDefault();
        setMessage('');
        if (!selectedUserId) {
            setMessage('Please select a user to delete.');
            return;
        }
        if (window.confirm('Are you sure you want to delete this user?')) {
            try {
                await userService.deleteUser(parseInt(selectedUserId));
                setMessage('User deleted successfully!');
                setSelectedUserId(''); // Clear selection
                fetchUsers(); // Refresh user list
            } catch (error) {
                setMessage('Error deleting user: ' + (error.response?.data?.error || error.message));
                console.error('Error deleting user:', error);
            }
        }
    };

    return (
        <div className="form-container">
            <h2>Delete User</h2>
            <div className="form-group">
                <label htmlFor="selectUserDelete">Select User:</label>
                <select id="selectUserDelete" value={selectedUserId} onChange={handleSelectChange}>
                    <option value="">-- Select a user --</option>
                    {users.map(user => (
                        <option key={user['id']} value={user['id']}>{user['username']} ({user['email']})</option>
                    ))}
                </select>
            </div>
            {selectedUserId && (
                <button onClick={handleDelete} className="submit-button delete-button">Delete User</button>
            )}
            {message && <p className="message">{message}</p>}
        </div>
    );
}

export default DeleteUser;