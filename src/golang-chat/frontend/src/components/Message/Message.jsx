import React from "react";
import { useState } from "react";
import "./Message.scss";

const Message = (props) => {
  const tmp = JSON.parse(props.message);
  const [message, setMessage] = useState(tmp);
  return <div className="Message">{message.body}</div>;
};

export default Message;
