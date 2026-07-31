use std::thread;
fn main(){
let vec4 = vec![0, 1, 2, 3];
let parallel_handle = thread::spawn(move || { // vec4がmoveする
println!("{:?}", vec4);
vec4 // vec4を返さないとこのスレッド内でvec4が解放される
});
// ここでスレッドから帰ってきたvec4の所有権を受け取る
match parallel_handle.join(){ // .join()でスレッドの終了を待つ
Ok(vec4) => println!("{:?}", vec4), // 無事、中身を取り出すことができる
Err(e) => println!("{:?}", e), // スレッド内でエラーが起こることもある
};
}
