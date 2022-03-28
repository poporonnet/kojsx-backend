import express, { Request, Response } from "express";
import * as db from "../../prisma/queries/main";

export const runsRouter = express.Router();
runsRouter.put("/:id", updateRun);
// ここの:idはhqが発行したid
async function updateRun(req: Request, res: Response) {
  let state: JudgeState = "AC";
  console.log(req.body);
  for (let i = 0; i < req.body.Status.length; i++) {
    if (req.body.Status[i] != "AC") {
      // 1つでもACでないものはACにならない
      state = req.body.Status[i].Status;
      if (state.length != 2) {
        state = "IE";
      }
    }
  }
  console.log("更新が走りました", state);
  const submission = await db.updateSubmissionState(
    req.params.id,
    Buffer.from(JSON.stringify(req.body.Status)).toString("base64"),
    state
  );
  res.json(submission);
}

type JudgeState = "CE" | "MLE" | "TLE" | "RE" | "OLE" | "IE" | "WA" | "AC";

/*

{
  "TaskID":"000-000",
  "Status":[
    {
      "TestID":"",
      "ExitStatus":0,
      "Duration":3118,
      "Status":"WA"
    },
    {
      "TestID":"",
      "ExitStatus":0,
      "Duration":3071,
      "Status":"WA"
    },
    {
      "TestID":"",
      "ExitStatus":0,
      "Duration":3069,
      "Status":"WA"
     }
   ]
}

*/
