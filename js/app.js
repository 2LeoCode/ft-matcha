"use strict";

/**
 * @typedef {{
 *   kind: "message";
 *   discussionId: number;
 *   senderId: number;
 *   content: string;
 * }} MessageNotif;
 *
 * @typedef {{
 *   kind: "kiss" | "punch" | "block";
 *   userId: number;
 * }} InteractionNotif;
 *
 * @typedef {{
 *   kind: "match";
 *   userId: number;
 * }} MatchNotif;
 *
 * @typedef {{
 *   kind: "block";
 *   userId: number;
 * }} BlockNotif;
 *
 * @typedef {MessageNotif | MatchNotif | BlockNotif} Notif;
 */

if (!window.WebSocket) {
  throw new Error("WebSockets are not supported by your browser");
}

/** @type {Notif[]} */
const notifs = [];
