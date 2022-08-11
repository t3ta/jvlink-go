meta:
  id: jv_ra_race
  endian: le
  encoding: sjis
seq:
  - id: records
    type: record
types:
  race_id:
    seq:
      - id: year
        type: str
        size: 4
      - id: month
        type: str
        size: 2
      - id: day
        type: str
        size: 2
      - id: jyo_code
        type: str
        size: 2
      - id: kaiji
        type: str
        size: 2
      - id: nichiji
        type: str
        size: 2
      - id: race_num
        type: str
        size: 2
  result:
    seq:
      - id: lap_time_strs
        type: str
        size: 3
        repeat: expr
        repeat-expr: 25
      - id: syogai_mile_time
        type: str
        size: 4
      - id: haron_time_s3_str
        type: str
        size: 3
      - id: haron_time_s4_str
        type: str
        size: 3
      - id: haron_time_l3_str
        type: str
        size: 3
      - id: haron_time_l4_str
        type: str
        size: 3
      - id: corner_1
        type: str
        size: 1
      - id: syukaisu_1
        type: str
        size: 1
      - id: jyuni_1
        type: str
        size: 70
      - id: corner_2
        type: str
        size: 1
      - id: syukaisu_2
        type: str
        size: 1
      - id: jyuni_2
        type: str
        size: 70
      - id: corner_3
        type: str
        size: 1
      - id: syukaisu_3
        type: str
        size: 1
      - id: jyuni_3
        type: str
        size: 70
      - id: corner_4_str
        type: str
        size: 1
      - id: syukaisu_4_str
        type: str
        size: 1
      - id: jyuni_4
        type: str
        size: 70
      - id: record_up_kubun
        size: 1
      - id: crlf
        size: 2
    instances:
      lap_times:
        type: lap_time(_index)
        repeat: expr
        repeat-expr: 25
      haron_time_s3:
        value: haron_time_s3_str.to_i
      haron_time_s4:
        value: haron_time_s4_str.to_i
      haron_time_l3:
        value: haron_time_l3_str.to_i
      haron_time_l4:
        value: haron_time_l4_str.to_i
  lap_time:
    params:
      - id: i
        type: s4
    instances:
      value:
        value: _parent.lap_time_strs[i].to_i
  record:
    seq:
      - id: record_spec
        type: str
        size: 2
      - id: data_kubun
        type: str
        size: 1
      - id: created_year
        type: str
        size: 4
      - id: created_month
        type: str
        size: 2
      - id: created_day
        type: str
        size: 2
      - id: race_id
        type: race_id
      - id: youbi_code
        type: str
        size: 1
      - id: toku_num
        type: str
        size: 4
      - id: hondai
        type: str
        size: 60
      - id: fukudai
        type: str
        size: 60
      - id: kakko
        type: str
        size: 60
      - id: hondai_eng
        type: str
        size: 120
      - id: fukudai_eng
        type: str
        size: 120
      - id: kakko_eng
        type: str
        size: 120
      - id: ryakusyo_10
        type: str
        size: 20
      - id: ryakusyo_6
        type: str
        size: 12
      - id: ryakusyo_3
        type: str
        size: 6
      - id: kubun
        type: str
        size: 1
      - id: nkai
        type: str
        size: 3
      - id: grade_code
        type: str
        size: 1
      - id: grade_code_before
        type: str
        size: 1
      - id: syubetu_code
        type: str
        size: 2
      - id: kigo_code
        type: str
        size: 3
      - id: jyuryo_code
        type: str
        size: 1
      - id: jyoken_code_2
        type: str
        size: 3
      - id: jyoken_code_3
        type: str
        size: 3
      - id: jyoken_code_4
        type: str
        size: 3
      - id: jyoken_code_5_years_old_and_over
        type: str
        size: 3
      - id: jyoken_code_youngest
        type: str
        size: 3
      - id: jyoken_name
        type: str
        size: 60
      - id: kyori_str
        type: str
        size: 4
      - id: kyori_before_str
        type: str
        size: 4
      - id: track_code
        type: str
        size: 2
      - id: track_code_before
        type: str
        size: 2
      - id: course_kubun_code
        type: str
        size: 2
      - id: course_kubun_code_before
        type: str
        size: 2
      - id: honsyokin
        type: str
        size: 8
        repeat: expr
        repeat-expr: 7
      - id: honsyokin_before
        type: str
        size: 8
        repeat: expr
        repeat-expr: 5
      - id: fukasyokin
        type: str
        size: 8
        repeat: expr
        repeat-expr: 5
      - id: fukasyokin_before
        type: str
        size: 8
        repeat: expr
        repeat-expr: 3
      - id: hasso_time
        type: str
        size: 4
      - id: hasso_time_before
        type: str
        size: 4
      - id: toroku_tosu
        type: str
        size: 2
      - id: syusso_tosu
        type: str
        size: 2
      - id: nyusen_tosu
        type: str
        size: 2
      - id: tenko_code
        type: str
        size: 1
      - id: siba_baba_code
        type: str
        size: 1
      - id: dirt_baba_code
        type: str
        size: 1
      - id: result
        type: result
        if: is_result
    instances:
      is_result:
        value: data_kubun.to_i > 2 and data_kubun.to_i < 8
      kyori:
        value: kyori_str.to_i
      kyori_before:
        value: kyori_before_str.to_i
