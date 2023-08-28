import React from 'react'
import { Link } from 'react-router-dom';
import './Navbar.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faHouse, faFilePen, faArrowRightFromBracket, faUser } from '@fortawesome/free-solid-svg-icons'

const Navbar = ({ isAuth }) => {
  return (
    <nav>
        <Link to='/'>
            <FontAwesomeIcon icon={faHouse} />ホーム
        </Link>
        {!isAuth ? (
            <Link to='login'>
                <FontAwesomeIcon icon={faArrowRightFromBracket} />ログイン
            </Link>
        ) : (
            <>
                <Link to='createPost'>
                    <FontAwesomeIcon icon={faFilePen} />記事投稿
                </Link>
                <Link to='#'>
                    <FontAwesomeIcon icon={faUser} />：{localStorage.getItem("userName")}
                </Link>
                <Link to='logout'>
                    <FontAwesomeIcon icon={faArrowRightFromBracket} />ログアウト
                </Link>
            </>
        )
        }
    </nav>
  )
}

export default Navbar