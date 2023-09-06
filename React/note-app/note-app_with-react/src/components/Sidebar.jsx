import React from 'react';
import './Sidebar.css';

const Sidebar = ({
  onAddNote,
  notes,
  sidebarActive,
  onSidebarIsActive,
  onDeleteNote,
  activeNote,
  setActiveNote
}) => {

    const sortedNotes = notes.sort((a, b) => b.modDate - a.modDate);

    return (sidebarActive ? (
      <div className='app-sidebar'>
        <div className="app-sidebar-header">
          <h1>ノート</h1>
          <button onClick={onAddNote}>追加</button>
        </div>
        <div className='nav-btn'>
          <button className='sidebar-button' onClick={onSidebarIsActive}>閉じる</button>
        </div>
        <div className="app-sidebar-notes">
          {sortedNotes.map((note) => (
            <div
              className={`app-sidebar-note ${note.id === activeNote && "active"}`}
              key={note.id}
              onClick={() => setActiveNote(note.id)}
            >
              <div className="note">
                <div className="sidebar-note-title">
                  <strong>{note.title}</strong>
                  <button onClick={() => onDeleteNote(note.id)}>削除</button>
                </div>
                <p>{note.content}</p>
              </div>
              <small>{new Date(note.modDate).toLocaleDateString("ja-jp")}</small>
            </div>
          ))}
        </div>
      </div>
    ) : (
      <div className='small-app-sidebar'>
        <button onClick={onSidebarIsActive}>開く</button>
        <button className='circle-border' onClick={() => {
          onAddNote(),
          onSidebarIsActive()
        }}>＋</button>
      </div>
    ))
}

export default Sidebar