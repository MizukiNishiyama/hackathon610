import './App.css';
import React, { useState, useEffect } from 'react';
// import { onAuthStateChanged } from "firebase/auth";
// import { fireAuth } from "./firebase";

type Channel = {
    id: string;
    name: string;
};

type Message = {
    messageid: string;
    content: string;
    userid: string;
    channelid: string;
};

type User = {
    id: string;
    name: string;
    email: string;
}

type Props = {
    activeChannel: string;
    setActiveChannel: React.Dispatch<React.SetStateAction<string>>;
}

function ShowChannelMessage(props:Props) {
    const {activeChannel, setActiveChannel} = props

    const [channels, setChannels] = useState<Channel[]>([]);
    const [messages, setMessages] = useState<Message[]>([]);


    useEffect(() => {
        const fetchChannels = async () => {
            const response = await fetch('http://localhost:8000/getchannels');
            const data = await response.json();
            setChannels(data);
        };
        fetchChannels();
    }, []);

    useEffect(() => {
        const fetchMessages = async () => {
            if (activeChannel === "") {
                setMessages([]);
                return;
            }
            const response = await fetch(`http://localhost:8000/message?channelid=${activeChannel}`);
            const data = await response.json();
            setMessages(data);
            // const response = await fetch(`http://localhost:8000/message?channelid=${activeChannel}`);
            // const data = await response.json();
            // setMessages(data);
        };

        fetchMessages();
    }, [activeChannel]);

    return (
        <div className="showmessages">
            <div className="channels">
                {channels.map(channel => (
                    <div
                        key={channel.id}
                        onClick={() => setActiveChannel(channel.id)}
                        className={activeChannel === channel.id ? 'active' : ''}
                    >
                        {channel.name}
                    </div>
                ))}
            </div>

            <div className="messages">
                {messages.map(message => (
                    <div key={message.messageid}>{message.content}</div>
                ))}
            </div>
        </div>
    );
}

function Sendmessage(props:Props) {
    // useEffect(() => {
    const {activeChannel, setActiveChannel} = props
    const [content, setContent] = useState("")
    const [userid ,setUserid] = useState("")
    
    const sendMessages = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        
        if (!content) {
            alert("メッセージを入力してください。");
            return;
        }
        
        try {
            const response = await fetch("http://localhost:8000/message", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    content: content,
                    channelid: activeChannel,
                    userid: userid,
                }),
                
            });

            const data = await response.json();
            console.log("success", data);
        } catch (error) {
            console.error("error:", error);
        }
    
    };
    
    // });    

    return(
        <div className="sendmessages">
                <form onSubmit={sendMessages}>
                    <label>
                        MESSAGE:
                        <input type="text" value={content} onChange={(e) => setContent(e.target.value)} />
                    </label>  
                    <label>
                        USERNAME:
                        <input type="text" value={userid} onChange={(e) => setUserid(e.target.value)} />
                    </label>  
                    <button type ="submit">SEND</button>
                </form>
            </div>
    )
}

function App() {
    const [activeChannel, setActiveChannel] = useState<string>("");
    return (
        <div>
            <ShowChannelMessage activeChannel={activeChannel} setActiveChannel={setActiveChannel}/>
            <Sendmessage activeChannel={activeChannel} setActiveChannel={setActiveChannel} />
        </div>
    )
}

export default App;


