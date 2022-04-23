import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { RpcError } from "@protobuf-ts/runtime-rpc";
import React, { useState } from "react";
import ReactDOM from "react-dom";
import { User } from "./gen/users/v1/users";
import { UserServiceClient } from "./gen/users/v1/users.client";

document.addEventListener("DOMContentLoaded", function () {
  ReactDOM.render(React.createElement(App), document.getElementById("root"));
});

const App: React.FC = () => {
  const transport = new GrpcWebFetchTransport({
    baseUrl: "/api",
  });
  const userService = new UserServiceClient(transport);
  return <UserForm userService={userService} />;
};

type UserFormProps = {
  userService: UserServiceClient;
};

const UserForm: React.FC<UserFormProps> = ({ userService }: UserFormProps) => {
  const [currentUser, setCurrentUser] = useState<User | undefined>();
  const [userName, setUsername] = useState<string>("");
  const [error, setError] = useState<RpcError | undefined>();

  if (error !== undefined) {
    return (
      <div>
        <h1>error: {error.message}</h1>
        <button onClick={() => setError(undefined)}>Clear error</button>
      </div>
    );
  }
  if (currentUser === undefined) {
    return (
      <form>
        <input
          value={userName}
          onChange={(event) => setUsername(event.target.value)}
          placeholder="Enter a username"
          type="text"
        ></input>
        <button
          onClick={async (event) => {
            event.preventDefault();
            try {
              const response = await userService.createUser({ name: userName });
              const user = response.response.user;
              if (user === undefined) {
                return;
              }
              setCurrentUser(user);
            } catch (error) {
              setError(error as RpcError);
            }
          }}
        >
          Create User
        </button>
      </form>
    );
  }
  return (
    <div>
      <h1>Your name: {currentUser.name}</h1>
      <h2>Your ID: {currentUser.id}</h2>
      <button onClick={() => setCurrentUser(undefined)}>Clear user</button>
    </div>
  );
};
