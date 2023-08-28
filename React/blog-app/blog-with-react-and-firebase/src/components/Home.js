import React, { useEffect, useState } from 'react'
import './Home.css';
import { collection, deleteDoc, doc, getDocs } from "firebase/firestore";
import { auth, db }  from "../firebase";

const Home = () => {
  const [postList, setPostList] = useState([]);

  useEffect(() => {
    const getPost = async () => {
      const data = await getDocs(collection(db, "posts"));
      setPostList(data.docs.map((doc) => {
        return {...doc.data(), id: doc.id};
      }));
    }
    getPost();
  }, []);

  const handleDelete = async (id) => {
    await deleteDoc(doc(db, 'posts', id));
    window.location.href = '/';
  }
  return (
    <div className='homePage'>
      {postList.map((post) => (
        <div className='postContents' key={post.id}>
          <div className='postHeader'>
            <h1>{post.title}</h1>
          </div>
          <div className='postTextContainer'>
            {post.postText}
          </div>
          <div className='nameAndDeleteButton'>
            <h3>{post.author.username}</h3>
            {post.author.id === auth.currentUser?.uid && (
              <button onClick={() => handleDelete(post.id)}>削除</button>
            )}
          </div>
        </div>
      ))}
    </div>
  )
}

export default Home