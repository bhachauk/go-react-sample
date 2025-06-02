import React, { useState, useEffect } from 'react';
import * as userService from '../services/userService';

function EditUser({ users, fetchUsers }) {
    const [selectedUserId, setSelectedUserId] = useState('');
    const [formData, setFormData] = useState({
        username: '',
        email: ''
    });
    const [message, setMessage] = useState('');

    useEffect(() => {
        if (selectedUserId) {
            const user = users.find(u => u.ID === parseInt(selectedUserId));
            if (user) {
                setFormData({ username: user['username'], email: user['email'] });
            }
        } else {
            setFormData({ username: '', email: '' }); // Clear form if no user selected
        }
    }, [selectedUserId, users]);

    const handleSelectChange = (e) => {
        setSelectedUserId(e.target.value);
        setMessage('');
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setMessage('');
        if (!selectedUserId) {
            setMessage('Please select a user to edit.');
            return;
        }
        try {
            await userService.updateUser(parseInt(selectedUserId), formData);
            setMessage('User updated successfully!');
            fetchUsers(); // Refresh user list
        } catch (error) {
            setMessage('Error updating user: ' + (error.response?.data?.error || error.message));
            console.error('Error updating user:', error);
        }
    };

    return (
        <div className="form-container">
            <h2>Edit User</h2>
            <div className="form-group">
                <label htmlFor="selectUser">Select User:</label>
                <select id="selectUser" value={selectedUserId} onChange={handleSelectChange}>
                    <option value="">-- Select a user --</option>
                    {users.map(user => (
                        <option key={user['id']} value={user['id']}>{user['username']} ({user['email']})</option>
                    ))}
                </select>
            </div>
            {selectedUserId && (
                <form onSubmit={handleSubmit}>
                    <div className="form-group">
                        <label htmlFor="editUsername">Username:</label>
                        <input
                            type="text"
                            id="editUsername"
                            name="username"
                            value={formData.username}
                            onChange={handleChange}
                            required
                        />
                    </div>
                    <div className="form-group">
                        <label htmlFor="editEmail">Email:</label>
                        <input
                            type="email"
                            id="editEmail"
                            name="email"
                            value={formData.email}
                            onChange={handleChange}
                            required
                        />
                    </div>
                    <button type="submit" className="submit-button">Update User</button>
                </form>
            )}
            {message && <p className="message">{message}</p>}
        </div>
    );
}

export default EditUser;