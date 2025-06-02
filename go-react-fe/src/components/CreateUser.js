import React, { useState } from 'react';
import * as userService from '../services/userService';

function CreateUser({ fetchUsers }) {
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: ''
    });
    const [message, setMessage] = useState('');

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setMessage('');
        try {
            await userService.createUser(formData);
            setMessage('User created successfully!');
            setFormData({ username: '', email: '', password: '' });
            fetchUsers(); // Refresh user list after creation
        } catch (error) {
            setMessage('Error creating user: ' + (error.response?.data?.error || error.message));
            console.error('Error creating user:', error);
        }
    };

    return (
        <div className="form-container">
            <h2>Create New User</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="username">Username:</label>
                    <input
                        type="text"
                        id="username"
                        name="username"
                        value={formData.username}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="email">Email:</label>
                    <input
                        type="email"
                        id="email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                    />
                </div>
                <button type="submit" className="submit-button">Create User</button>
            </form>
            {message && <p className="message">{message}</p>}
        </div>
    );
}

export default CreateUser;