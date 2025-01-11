export type MessageNotif = {
  kind: "message";
  userId: number;
  discussionId: number;
  content: string;
};

export type InteractionNotif = {
  kind: "kiss" | "punch" | "block";
  userId: number;
};

export type Notif = MessageNotif | InteractionNotif;

export type NotifHandler = (notif: Notif) => void;

export type OnNotifOptions = {
  once: boolean;
};

if (!window.WebSocket)
  throw new Error("Your browser does not support WebSockets");

const socket = new WebSocket("ws://localhost:3000");

const notifEventTarget = new EventTarget();

socket.addEventListener("message", (e) => {
  const notif: Notif = e.data;
  const notifEvent = new CustomEvent<Notif>(notif.kind, {
    detail: notif,
  });
  notifEventTarget.dispatchEvent(notifEvent);
});

const onNotifDefaultOptions: OnNotifOptions = {
  once: false,
};

export function onNotif<Kind extends Notif["kind"]>(
  kind: Kind,
  handler: (notif: Extract<Notif, { kind: Kind }>) => void,
  options: OnNotifOptions = onNotifDefaultOptions,
) {
  const wrapper = (e: Event) => {
    const notifEvent = e as CustomEvent<Extract<Notif, { kind: Kind }>>;
    handler(notifEvent.detail);
  };
  notifEventTarget.addEventListener(kind, wrapper, options);
  return () => notifEventTarget.removeEventListener(kind, wrapper);
}
