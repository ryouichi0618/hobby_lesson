import React from 'react'
import { Link } from 'react-router-dom';
import './Navbar.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faHouse, faFilePen, faArrowRightFromBracket, faUser } from '@fortawesome/free-solid-svg-icons'
import { auth } from '../firebase';

const Navbar = ({ isAuth }) => {
  return (
    <nav>
        <Link to='/'>
            <FontAwesomeIcon icon={faHouse} />ホーム
        </Link>
        <Link to='createPost'>
            <FontAwesomeIcon icon={faFilePen} />記事投稿
        </Link>
        {!isAuth ? (
            <Link to='login'>
                <FontAwesomeIcon icon={faArrowRightFromBracket} />ログイン
            </Link>
        ) : (
            <Link to='logout'>
                <FontAwesomeIcon icon={faArrowRightFromBracket} />ログアウト
            </Link>
        )
        }
        {isAuth && (
        <Link to='#'>
            <FontAwesomeIcon icon={faUser} />：{auth.currentUser.displayName}
        </Link>
        )}
    </nav>
  )
}

export default Navbar