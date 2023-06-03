import './App.css';
import React, { useState, useEffect } from 'react';

//authentification
import { onAuthStateChanged } from "firebase/auth";
import { fireAuth } from "./firebase";
import { signInWithPopup, GoogleAuthProvider, signOut } from "firebase/auth";

// //routing
import { BrowserRouter, Route} from 'react-router-dom';

export const LoginForm: React.FC = () => {
    /**
     * googleでログインする
     */
    const signInWithGoogle = (): void => {
      // Google認証プロバイダを利用する
      const provider = new GoogleAuthProvider();
  
      // ログイン用のポップアップを表示
      signInWithPopup(fireAuth, provider)
        .then(res => {
            const user =res.user;
            alert("ログインユーザー: " + user.displayName);
            fetch("http://localhost:8000/user", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    name: user.displayName,
                    email: user.email
                })
            })    
        .then(response => response.json())
        .then(data => console.log(data))
        .catch((error) => {
            console.error('Error:', error);
        });
        })
        .catch(err => {
            const errorMessage =err.message;
            alert(errorMessage);
        })
    };
  
    /**
     * ログアウトする
     */
    const signOutWithGoogle = (): void => {
        signOut(fireAuth).then(() => {
          alert("ログアウトしました");
        }).catch(err => {
          alert(err);
        });
      };
    
  
    return (
      <div className = "loginform" >
        <button onClick={signInWithGoogle}>
          Googleでログイン
        </button>
        <button onClick={signOutWithGoogle}>
        ログアウト
        </button>
      </div>
    );
  };


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
            
        };

        fetchMessages();
    }, [activeChannel]);
    

    return (
        <div className="showmessages">
            <div className="channels">
                <h1>Channel</h1>
                {channels.map(channel => (
                    <div
                        key={channel.id}
                        onClick={() =>  setActiveChannel(channel.id)}
                        className={activeChannel === channel.id ? 'active' : ''}
                    >
                        {channel.name}
                    </div>
                ))}
            </div>
            <h1>Talk</h1>        
            <div className="messages">
                {messages.map(message => (
                    <div key={message.messageid}>{message.content}</div>
                ))}
            </div>
        </div>
    );
}

// function getUserinfo () {
    

// }

function Sendmessage(props:Props) {
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
            const user = fireAuth.currentUser
            if (!user) {
                alert("ユーザーがログインしていません。");
                return;
            }
            const response = await fetch("http://localhost:8000/message", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    content: content,
                    channelid: activeChannel,
                    userid: user.displayName,
                }),
                
            });

            const data = await response.json();
            console.log("success", data);
        } catch (error) {
            console.error("error:", error);
        }
    
    };
 
    return(
        <div className="sendmessages">
                <form onSubmit={sendMessages}>
                    <h1>
                        MESSAGE 
                        <input type="text" value={content} onChange={(e) => setContent(e.target.value)} />
                    </h1>  
                    <button type ="submit">SEND</button>
                </form>
            </div>
    )
}

function LoginState() {
    const [loginUser, setLoginUser] = useState(fireAuth.currentUser);
  
  // ログイン状態を監視して、stateをリアルタイムで更新する
    onAuthStateChanged(fireAuth, user => {
        setLoginUser(user);
    });
    return(
        <div>
            <LoginForm />
            {loginUser ? <App /> : null}
        </div>
        
    );
}


function App() {
    const [activeChannel, setActiveChannel] = useState<string>("");

    return (
        // <BrowserRouter>
                <div>
                    <LoginForm /> 
                    <ShowChannelMessage activeChannel={activeChannel} setActiveChannel={setActiveChannel}/>
                    <Sendmessage activeChannel={activeChannel} setActiveChannel={setActiveChannel} />     
                                
                </div>
        // </BrowserRouter>
    );
}

export default App;


