import './App.css';
import React, { useState, useEffect } from 'react';

//authentification
import { onAuthStateChanged } from "firebase/auth";
import { fireAuth } from "./firebase";
import { signInWithPopup, GoogleAuthProvider, signOut } from "firebase/auth";

//material UI

import Button from '@mui/material/Button';
import SendIcon from '@mui/icons-material/Send';
import DeleteIcon from '@mui/icons-material/Delete';
import Box from '@mui/material/Box';
import IconButton from '@mui/material/IconButton';
import AppBar from '@mui/material/AppBar';
import EditIcon from '@mui/icons-material/Edit';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import AddIcon from '@mui/icons-material/Add';


export const LoginForm: React.FC = () => {
    
    const signInWithGoogle = (): void => {
      // Google認証プロバイダを利用する
      const provider = new GoogleAuthProvider();
  
      signInWithPopup(fireAuth, provider)
        .then(res => {
            const user =res.user;
            alert("ログインユーザー: " + user.displayName);
            // fetch("http://localhost:3000/user", {
            fetch("https://uttc-bapgglyr6q-uc.a.run.app/user", {
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
    
    // fucntion loginname() {
    //     const [user, setUser]=useState(null);
    //     useEffect(() => {
    //         fireAuth.onAuthStateChanged(function (user) {
    //             if (user) {
    //                 setUser(user);
    //             } else {
    //                 setUser(null);
    //             }
    //         });
    //     },[]);
    // }
  
    return (
      <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h5" component="div" sx={{ flexGrow: 1 }}>
            Message Application
          </Typography>
          <Button color="inherit"  onClick={signInWithGoogle} sx={{ margin:"3px", bgcolor:"#0000ff4a"}}>Login</Button>
          <Button color="inherit"  onClick={signOutWithGoogle} sx={{ margin:"3px", bgcolor:"#0000ff4a"}}>Logout</Button>
          <Typography variant="h6" component="div" sx={{marginLeft:"20px" ,textAlign:"right", fontSize:"16px"}}>
            {fireAuth.currentUser ? fireAuth.currentUser.displayName:""}
          </Typography>
        </Toolbar>
      </AppBar>
    </Box>
    );
  };


type Channel = {
    id: string;
    name: string;
};

type Message = {
    id: string;
    content: string;
    userid: string;
    channelid: string;
    time: string;
};

type User = {
    name: string;
    email: string;
}

type Props = {
    activeChannel: string;
    setActiveChannel: React.Dispatch<React.SetStateAction<string>>;
    refreshMessages: boolean;
    setRefreshMessages: React.Dispatch<React.SetStateAction<boolean>>;
}

type EditableMessageProps = {
    message: Message;
    deleteMessage: (messageId: string) => void;
    editMessage: (messageId: string, messageContent: string) => void;
};

const EditableMessage: React.FC<EditableMessageProps> = ({ message, deleteMessage, editMessage }) => {
    const [isEditing, setIsEditing] = useState(false);
    const [editContent, setEditContent] = useState(message.content);
    return (
        <div key={message.time}>
            <span className="user-name">{message.userid}</span>
            <div className="content-time">
                <span className="message-content">{message.content}</span>
                <span className="message-time">{message.time}</span>
                <Box justifyContent="space-between">                
                <IconButton color="error" aria-label="delete" onClick={() => deleteMessage(message.id)}><DeleteIcon /></IconButton>
                {isEditing ? (
                    <form onSubmit={(event) => {
                        event.preventDefault();
                        editMessage(message.id, editContent);
                        setIsEditing(false);  
                    }}>
                        <textarea 
                            value={editContent}
                            onChange={event => setEditContent(event.target.value)}
                        />
                        <IconButton color="primary" type="submit" aria-label="send" size="small"><SendIcon /></IconButton>
                    </form>
                ) : (
                    <IconButton color="success" aria-label="edit" onClick={() => setIsEditing(true)} size="small"><EditIcon /></IconButton>
                )}
                </Box> 
                
            </div>
        </div>
    );
};


function ShowChannelMessage(props:Props) {
    const {activeChannel, setActiveChannel, refreshMessages, setRefreshMessages} = props

    const [channels, setChannels] = useState<Channel[]>([]);
    const [messages, setMessages] = useState<Message[]>([]);
    const [isEditing, setIsEditing] = useState(false);
    const [editchannelname, setEditchannelname] = useState("");
    const [channelname, setName] = useState("")
    
    useEffect(() => {
        const fetchChannels = async () => {
            const response = await fetch('https://uttc-bapgglyr6q-uc.a.run.app/getchannels');
            const data = await response.json();
            console.log("表示するチャンネル",data);
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
            const response = await fetch(`https://uttc-bapgglyr6q-uc.a.run.app/message?channelid=${activeChannel}`);
            const data = await response.json();
            console.log("表示するメッセージ",data);
            setMessages(data);
            
        };
        fetchMessages();
    }, [activeChannel, refreshMessages]);
    
    const editChannel = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        // setName=("")
        try {
            const response = await fetch("https://uttc-bapgglyr6q-uc.a.run.app/channel", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    Name: editchannelname,
                }),
                
            });
            const data = await response.json();
            console.log("success", data);
            setEditchannelname("");
            alert("チャンネルを追加しました。リロードしてください。");
        } catch (error) {
            console.error("error:", error);
        }
    }

    async function deleteMessage(messageId: string) {
        console.log("削除するメッセージ", messageId)
        try {
            const response = await fetch(`https://uttc-bapgglyr6q-uc.a.run.app/delete_message?id=${messageId}`, {
                method: 'DELETE',
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            setRefreshMessages(!refreshMessages);
            alert("メッセージを削除しました。");
        } catch (error) {
            console.error("An error occurred while deleting the message:", error);
        }
    }
    
    async function EditMessage(messageId: string, messageContent: string) {
        console.log("編集するメッセージ", messageId, messageContent)
        try {
            const response = await fetch(`https://uttc-bapgglyr6q-uc.a.run.app/edit?id=${messageId}&content=${messageContent}`, {
                method: 'POST',
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            setRefreshMessages(!refreshMessages);
            alert("メッセージを編集しました。");
        } catch (error) {
            console.error("An error occurred while deleting the message:", error);
        }

    }
    
    
    return (
        <div className="showmessages">
            <div className="channels">
            <div className='channel_label'>Channels</div>
            <div className='channel_list'>
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
            <div className='addchannel'>   
                {isEditing ? (
                    <form onSubmit={(event) => {
                        event.preventDefault();
                        editChannel(event);
                        setIsEditing(false);  
                    }}>
                        <textarea
                            placeholder="Enter the name of new channel..." 
                            value={editchannelname}
                            onChange={event => setEditchannelname(event.target.value)}
                        />
                        <IconButton color="primary" type="submit" aria-label="send" size="small" ><SendIcon /></IconButton>
                    </form>
                ) : (
                    // <IconButton color="primary" aria-label="add" onClick={() => setIsEditing(true)} size="small"><AddIcon /></IconButton>
                    <Button size="large" color="primary" variant="outlined" endIcon={<AddIcon />} onClick={() => setIsEditing(true)}>Add Channel</Button>
                )}
            </div>    
            </div>
                
            <div className="messages">
                {messages.map(message => (
                    <EditableMessage
                        key={message.id}
                        message={message}
                        deleteMessage={deleteMessage}
                        editMessage={EditMessage}
                    />
                ))}
            </div>
        </div>
    );
}



function Sendmessage(props:Props) {
    const {activeChannel, refreshMessages, setRefreshMessages} = props
    const [content, setContent] = useState("")
    
    
    
    const sendMessages = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log(content);
        setContent("")

        if (!content) {
            alert("メッセージを入力してください。");
            return;
        }
        if (activeChannel === "") {
            alert("チャンネルを選択してください。");
            return;
        }
        try {
            const user = fireAuth.currentUser
            if (!user) {
                alert("ログインしてください");
                return;
            }
            const response = await fetch("https://uttc-bapgglyr6q-uc.a.run.app/message", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    content: content,
                    channelid: activeChannel,
                    userid: user.displayName,
                    time: ` (${new Date().getMonth()+1}/${new Date().getDate()} ${new Date().getHours()}:${new Date().getMinutes()})`,
                }),
                
            });

            const data = await response.json();
            console.log("success", data);
            setRefreshMessages(!refreshMessages)
        } catch (error) {
            console.error("error:", error);
        }
    
    };
 
    return(
        <div className="sendmessages">
                <form onSubmit={sendMessages}>
                    <textarea placeholder="Type here!" value={content} onChange={(e) => setContent(e.target.value)} />
                    
                    <Button type ="submit" variant="contained" endIcon={<SendIcon />} >SEND</Button>
                </form>
            </div>
    )
}

// function LoginState() {
//     const [loginUser, setLoginUser] = useState(fireAuth.currentUser);
  
//   // ログイン状態を監視して、stateをリアルタイムで更新する
//     onAuthStateChanged(fireAuth, user => {
//         setLoginUser(user);
//     });
//     return(
//         <div>
//             <LoginForm />
//             {loginUser ? <App /> : null}
//         </div>
        
//     );
// }


function App() {
    const [activeChannel, setActiveChannel] = useState<string>("");
    const [refreshMessages, setRefreshMessages] = useState<boolean>(false);

    // useEffect(() => {
    //     const interval =setInterval(() => {
    //         window.location.reload();
    //     }, 5000);
    //     return () => clearInterval(interval);
    // }, []);

    return (
        // <BrowserRouter>
                <div className="App">
                    <LoginForm /> 
                    <ShowChannelMessage activeChannel={activeChannel} setActiveChannel={setActiveChannel} setRefreshMessages={setRefreshMessages} refreshMessages={refreshMessages}/>
                    <Sendmessage activeChannel={activeChannel} setActiveChannel={setActiveChannel} setRefreshMessages={setRefreshMessages} refreshMessages={refreshMessages}/>     
                                
                </div>
        // </BrowserRouter>
    );
}

export default App;


