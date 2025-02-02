import { create } from "zustand";

type SessionStore = {
  userId: string;
  username: string;
};

export const useSessionStore = create<SessionStore>((set) => {
  const userId = crypto.randomUUID();
  const username = randomUsername();

  return { userId, username };
});

function randomUsername() {
  const usernames = [
    "shapemouthguard",
    "frazzledwhite",
    "raggedheritage",
    "replicamindless",
    "movekapi",
    "uprightcare",
    "consultingsurely",
    "isleunfriendly",
    "unitquiver",
    "handmadeenigmatic",
    "soresailor",
    "everywhereinertia",
    "cancertour",
    "controlamphora",
    "pressurejackstay",
    "popularsearch",
    "suchoranges",
    "convertinghomework",
    "runslurp",
    "artichokestreamer",
    "nevercrummy",
    "producecircle",
    "thawdrafter",
    "pestradio",
    "fiftherase",
    "decentmisguided",
    "greetpyramid",
    "googlehibiscus",
    "husbandwindbound",
    "abhorrentsandwich",
    "noseregard",
    "comedygoujon",
    "notionendurable",
    "energeticslice",
    "cameramojang",
    "bullocksbucket",
    "dreamfiery",
    "gendermessage",
    "conceitedmainsail",
    "announcepiercer",
    "whisperwage",
    "latersubtle",
    "noticedinner",
    "poplardepending",
    "bootsadmission",
    "axelfrequency",
    "therapistfearless",
    "beaconbombast",
    "carefreedisturbed",
    "saviorabaft",
  ];

  const randomId = Math.floor((Math.random() * 10) % usernames.length);

  return usernames[randomId];
}
