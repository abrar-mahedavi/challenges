import java.util.*;

class Program {

  public String tournamentWinner(
    ArrayList<ArrayList<String>> competitions, ArrayList<Integer> results) {
    // Write your code here.
    // return "";
    Map<String,Integer> competitorsScore = new HashMap<>();
    for(int i = 0;i<results.size();i++){
      if(results.get(i)==0){
        if(competitorsScore.containsKey(competitions.get(i).get(1))){
          Integer score = competitorsScore.get(competitions.get(i).get(1));
          competitorsScore.put(competitions.get(i).get(1),score+3);
        }else{
          competitorsScore.put(competitions.get(i).get(1),3);
        }
      }else{
        if(competitorsScore.containsKey(competitions.get(i).get(0))){
          Integer score = competitorsScore.get(competitions.get(i).get(0));
          competitorsScore.put(competitions.get(i).get(0),score+3);
        }else{
          competitorsScore.put(competitions.get(i).get(0),3);
        }
      }
    }
    int maxGameWinner = 0;
    String winner = "";
    for(Map.Entry<String,Integer> comp : competitorsScore.entrySet()){
      if (comp.getValue() > maxGameWinner){
        maxGameWinner = comp.getValue();
        winner = comp.getKey();
      }
    }
    return winner;
  }
}
