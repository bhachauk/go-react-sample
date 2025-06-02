import React from 'react';

function UserTable({ users }) {
    return (
        <div className="user-table-container">
            <h2>All Users</h2>
            {users.length === 0 ? (
                <p>No users found.</p>
            ) : (
                <table>
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>Username</th>
                        <th>Email</th>
                        <th>Created At</th>
                        <th>Updated At</th>
                    </tr>
                    </thead>
                    <tbody>
                    {users.map(user => (
                        <tr key={user['id']}>
                            <td>{user['id']}</td>
                            <td>{user['username']}</td>
                            <td>{user.Email}</td>
                            <td>{new Date(user.CreatedAt).toLocaleDateString()}</td>
                            <td>{new Date(user.UpdatedAt).toLocaleDateString()}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
}

export default UserTable;